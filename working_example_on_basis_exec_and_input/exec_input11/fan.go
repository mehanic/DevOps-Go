package main

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"log"
	"strings"
	"sync"
)

func SourceLineWords(ctx context.Context, r io.ReadCloser) <-chan []string {
	ch := make(chan []string)
	go func() {
		defer func() { r.Close(); close(ch) }()
		b := bytes.Buffer{}
		s := bufio.NewScanner(r)
		for s.Scan() {
			b.Reset()
			b.Write(s.Bytes())
			words := []string{}
			w := bufio.NewScanner(&b)
			w.Split(bufio.ScanWords)
			for w.Scan() {
				words = append(words, w.Text())
			}
			select {
			case <-ctx.Done():
				return
			case ch <- words:
			}
		}
	}()
	return ch
}

func WordOccurrence(ctx context.Context, src <-chan []string) <-chan map[string]int {
	ch := make(chan map[string]int)
	go func() {
		defer close(ch)
		for v := range src {
			count := make(map[string]int)
			for _, s := range v {
				count[s]++
			}
			select {
			case <-ctx.Done():
				return
			case ch <- count:
			}
		}
	}()
	return ch
}

func MergeCounts(ctx context.Context, src ...<-chan map[string]int) map[string]int {
	count := make(map[string]int)
	wg := sync.WaitGroup{}
	merge := make(chan map[string]int)
	wg.Add(len(src))
	go func() {
		wg.Wait()
		close(merge)
	}()
	for _, ch := range src {
		go func(ch <-chan map[string]int) {
			defer wg.Done()
			for v := range ch {
				select {
				case <-ctx.Done():
					return
				case merge <- v:
				}
			}
		}(ch)
	}
	for {
		select {
		case <-ctx.Done():
			return count
		case c, ok := <-merge:
			if !ok {
				return count
			}
			for k, v := range c {
				count[k] += v
			}
		}
	}
}

func main() {
	ctx, canc := context.WithCancel(context.Background())
	defer canc()
	src := SourceLineWords(ctx, ioutil.NopCloser(strings.NewReader(cantoUno)))
	count1, count2 := WordOccurrence(ctx, src), WordOccurrence(ctx, src)
	final := MergeCounts(ctx, count1, count2)
	log.Println(final)
}

const cantoUno = `Nel mezzo del cammin di nostra vita
mi ritrovai per una selva oscura
ché la diritta via era smarrita.
Ahi quanto a dir qual era è cosa dura
esta selva selvaggia e aspra e forte
che nel pensier rinova la paura!
Tant’è amara che poco è più morte;
ma per trattar del ben ch’i’ vi trovai,
dirò de l’altre cose ch’i’ v’ho scorte.
Io non so ben ridir com’i’ v’intrai,
tant’era pien di sonno a quel punto
che la verace via abbandonai.
Ma poi ch’i’ fui al piè d’un colle giunto,
là dove terminava quella valle
che m’avea di paura il cor compunto,
guardai in alto, e vidi le sue spalle
vestite già de’ raggi del pianeta
che mena dritto altrui per ogne calle.
Allor fu la paura un poco queta
che nel lago del cor m’era durata
la notte ch’i’ passai con tanta pieta.
E come quei che con lena affannata
uscito fuor del pelago a la riva
si volge a l’acqua perigliosa e guata,
così l’animo mio, ch’ancor fuggiva,
si volse a retro a rimirar lo passo
che non lasciò già mai persona viva.
Poi ch’èi posato un poco il corpo lasso,
ripresi via per la piaggia diserta,
sì che ’l piè fermo sempre era ’l più basso.
Ed ecco, quasi al cominciar de l’erta,
una lonza leggera e presta molto,
che di pel macolato era coverta;
e non mi si partia dinanzi al volto,
anzi ’mpediva tanto il mio cammino,
ch’i’ fui per ritornar più volte vòlto.
Temp’era dal principio del mattino,
e ’l sol montava ’n sù con quelle stelle
ch’eran con lui quando l’amor divino
mosse di prima quelle cose belle;
sì ch’a bene sperar m’era cagione
di quella fiera a la gaetta pelle
l’ora del tempo e la dolce stagione;
ma non sì che paura non mi desse
la vista che m’apparve d’un leone.
Questi parea che contra me venisse
con la test’alta e con rabbiosa fame,
sì che parea che l’aere ne tremesse.
Ed una lupa, che di tutte brame
sembiava carca ne la sua magrezza,
e molte genti fé già viver grame,
questa mi porse tanto di gravezza
con la paura ch’uscia di sua vista,
ch’io perdei la speranza de l’altezza.
E qual è quei che volontieri acquista,
e giugne ’l tempo che perder lo face,
che ’n tutt’i suoi pensier piange e s’attrista;
tal mi fece la bestia sanza pace,
che, venendomi ’ncontro, a poco a poco
mi ripigneva là dove ’l sol tace.
Mentre ch’i’ rovinava in basso loco,
dinanzi a li occhi mi si fu offerto
chi per lungo silenzio parea fioco.
Quando vidi costui nel gran diserto,
«Miserere di me», gridai a lui,
«qual che tu sii, od ombra od omo certo!».
Rispuosemi: «Non omo, omo già fui,
e li parenti miei furon lombardi,
mantoani per patria ambedui.
Nacqui sub Iulio, ancor che fosse tardi,
e vissi a Roma sotto ’l buono Augusto
nel tempo de li dèi falsi e bugiardi.
Poeta fui, e cantai di quel giusto
figliuol d’Anchise che venne di Troia,
poi che ’l superbo Ilión fu combusto.
Ma tu perché ritorni a tanta noia?
perché non sali il dilettoso monte
ch’è principio e cagion di tutta gioia?».
«Or se’ tu quel Virgilio e quella fonte
che spandi di parlar sì largo fiume?»,
rispuos’io lui con vergognosa fronte.
«O de li altri poeti onore e lume
vagliami ’l lungo studio e ’l grande amore
che m’ha fatto cercar lo tuo volume.
Tu se’ lo mio maestro e ’l mio autore;
tu se’ solo colui da cu’ io tolsi
lo bello stilo che m’ha fatto onore.
Vedi la bestia per cu’ io mi volsi:
aiutami da lei, famoso saggio,
ch’ella mi fa tremar le vene e i polsi».
«A te convien tenere altro viaggio»,
rispuose poi che lagrimar mi vide,
«se vuo’ campar d’esto loco selvaggio:
ché questa bestia, per la qual tu gride,
non lascia altrui passar per la sua via,
ma tanto lo ’mpedisce che l’uccide;
e ha natura sì malvagia e ria,
che mai non empie la bramosa voglia,
e dopo ’l pasto ha più fame che pria.
Molti son li animali a cui s’ammoglia,
e più saranno ancora, infin che ’l veltro
verrà, che la farà morir con doglia.
Questi non ciberà terra né peltro,
ma sapienza, amore e virtute,
e sua nazion sarà tra feltro e feltro.
Di quella umile Italia fia salute
per cui morì la vergine Cammilla,
Eurialo e Turno e Niso di ferute.
Questi la caccerà per ogne villa,
fin che l’avrà rimessa ne lo ’nferno,
là onde ’nvidia prima dipartilla.
Ond’io per lo tuo me’ penso e discerno
che tu mi segui, e io sarò tua guida,
e trarrotti di qui per loco etterno,
ove udirai le disperate strida,
vedrai li antichi spiriti dolenti,
ch’a la seconda morte ciascun grida;
e vederai color che son contenti
nel foco, perché speran di venire
quando che sia a le beate genti.
A le quai poi se tu vorrai salire,
anima fia a ciò più di me degna:
con lei ti lascerò nel mio partire;
ché quello imperador che là sù regna,
perch’i’ fu’ ribellante a la sua legge,
non vuol che ’n sua città per me si vegna.
In tutte parti impera e quivi regge;
quivi è la sua città e l’alto seggio:
oh felice colui cu’ ivi elegge!».
E io a lui: «Poeta, io ti richeggio
per quello Dio che tu non conoscesti,
acciò ch’io fugga questo male e peggio,
che tu mi meni là dov’or dicesti,
sì ch’io veggia la porta di san Pietro
e color cui tu fai cotanto mesti».
Allor si mosse, e io li tenni dietro.`



//2025/02/14 20:15:03 map[A:1 Ahi:1 Allor:2 Augusto:1 Cammilla,:1 Di:1 Dio:1 E:3 Ed:2 Eurialo:1 Ilión:1 In:1 Io:1 Italia:1 Iulio,:1 Ma:2 Mentre:1 Molti:1 Nacqui:1 Nel:1 Niso:1 Ond’io:1 Pietro:1 Poeta:1 Poi:1 Quando:1 Questi:3 Rispuosemi::1 Roma:1 Tant’è:1 Temp’era:1 Troia,:1 Tu:1 Turno:1 Vedi:1 Virgilio:1 a:18 abbandonai.:1 acciò:1 acquista,:1 affannata:1 aiutami:1 al:3 alto,:1 altri:1 altro:1 altrui:2 amara:1 ambedui.:1 amore:2 ancor:1 ancora,:1 anima:1 animali:1 antichi:1 anzi:1 aspra:1 autore;:1 basso:1 basso.:1 beate:1 belle;:1 bello:1 ben:2 bene:1 bestia:2 bestia,:1 brame:1 bramosa:1 bugiardi.:1 buono:1 caccerà:1 cagion:1 cagione:1 calle.:1 cammin:1 cammino,:1 campar:1 cantai:1 carca:1 cercar:1 certo!».:1 che:40 che,:1 chi:1 ché:3 ch’a:2 ch’ancor:1 ch’ella:1 ch’eran:1 ch’io:3 ch’i’:6 ch’uscia:1 ch’è:1 ch’èi:1 ciascun:1 ciberà:1 città:2 ciò:1 colle:1 color:2 colui:2 combusto.:1 come:1 cominciar:1 compunto,:1 com’i’:1 con:10 conoscesti,:1 contenti:1 contra:1 convien:1 cor:2 corpo:1 cosa:1 cose:2 costui:1 così:1 cotanto:1 coverta;:1 cui:3 cu’:3 da:2 dal:1 de:5 degna::1 del:7 desse:1 de’:1 di:19 dicesti,:1 dietro.:1 dilettoso:1 dinanzi:2 dipartilla.:1 dir:1 diritta:1 dirò:1 discerno:1 diserta,:1 diserto,:1 disperate:1 divino:1 doglia.:1 dolce:1 dolenti,:1 dopo:1 dove:2 dov’or:1 dritto:1 dura:1 durata:1 dèi:1 d’Anchise:1 d’esto:1 d’un:2 e:40 ecco,:1 elegge!».:1 empie:1 era:4 esta:1 etterno,:1 fa:1 face,:1 fai:1 falsi:1 fame:1 fame,:1 famoso:1 farà:1 fatto:2 fece:1 felice:1 feltro:1 feltro.:1 fermo:1 ferute.:1 fia:2 fiera:1 figliuol:1 fin:1 fioco.:1 fiume?»,:1 foco,:1 fonte:1 forte:1 fosse:1 fronte.:1 fu:3 fugga:1 fuggiva,:1 fui:2 fui,:2 fuor:1 furon:1 fu’:1 fé:1 gaetta:1 genti:1 genti.:1 gioia?».:1 giugne:1 giunto,:1 giusto:1 già:4 grame,:1 gran:1 grande:1 gravezza:1 grida;:1 gridai:1 gride,:1 guardai:1 guata,:1 guida,:1 ha:2 i:1 il:4 impera:1 imperador:1 in:2 infin:1 io:6 ivi:1 la:26 lago:1 lagrimar:1 largo:1 lascerò:1 lascia:1 lasciò:1 lasso,:1 le:5 legge,:1 leggera:1 lei:1 lei,:1 lena:1 leone.:1 li:7 lo:8 loco:2 loco,:1 lombardi,:1 lonza:1 lui:2 lui,:1 lui::1 lume:1 lungo:2 lupa,:1 là:5 l’acqua:1 l’aere:1 l’altezza.:1 l’alto:1 l’altre:1 l’amor:1 l’animo:1 l’avrà:1 l’erta,:1 l’ora:1 l’uccide;:1 ma:4 macolato:1 maestro:1 magrezza,:1 mai:2 male:1 malvagia:1 mantoani:1 mattino,:1 me:3 mena:1 meni:1 mesti».:1 mezzo:1 me»,:1 me’:1 mi:12 miei:1 mio:4 mio,:1 molte:1 molto,:1 montava:1 monte:1 morir:1 morte:1 morte;:1 morì:1 mosse:1 mosse,:1 m’apparve:1 m’avea:1 m’era:2 m’ha:2 natura:1 nazion:1 ne:3 nel:6 noia?:1 non:11 nostra:1 notte:1 né:1 occhi:1 od:2 offerto:1 ogne:2 oh:1 ombra:1 omo:2 omo,:1 onde:1 onore:1 onore.:1 oscura:1 ove:1 pace,:1 parea:3 parenti:1 parlar:1 parti:1 partia:1 partire;:1 passai:1 passar:1 passo:1 pasto:1 patria:1 paura:4 paura!:1 peggio,:1 pel:1 pelago:1 pelle:1 peltro,:1 pensier:2 penso:1 per:16 perché:3 perch’i’:1 perdei:1 perder:1 perigliosa:1 persona:1 piaggia:1 pianeta:1 piange:1 pien:1 pieta.:1 piè:2 più:6 poco:5 poeti:1 poi:4 polsi».:1 porse:1 porta:1 posato:1 presta:1 pria.:1 prima:2 principio:2 punto:1 quai:1 qual:3 quando:2 quanto:1 quasi:1 quei:2 quel:3 quella:4 quelle:2 quello:2 questa:2 questo:1 queta:1 qui:1 quivi:2 rabbiosa:1 raggi:1 regge;:1 regna,:1 retro:1 ria,:1 ribellante:1 richeggio:1 ridir:1 rimessa:1 rimirar:1 rinova:1 ripigneva:1 ripresi:1 rispuose:1 rispuos’io:1 ritornar:1 ritorni:1 ritrovai:1 riva:1 rovinava:1 saggio,:1 sali:1 salire,:1 salute:1 san:1 sanza:1 sapienza,:1 saranno:1 sarà:1 sarò:1 scorte.:1 se:1 seconda:1 seggio::1 segui,:1 selva:2 selvaggia:1 selvaggio::1 sembiava:1 sempre:1 se’:3 si:6 sia:1 sii,:1 silenzio:1 smarrita.:1 so:1 sol:2 solo:1 son:2 sonno:1 sotto:1 spalle:1 spandi:1 speran:1 speranza:1 sperar:1 spiriti:1 stagione;:1 stelle:1 stilo:1 strida,:1 studio:1 sua:7 sub:1 sue:1 suoi:1 superbo:1 sì:7 sù:2 s’ammoglia,:1 s’attrista;:1 tace.:1 tal:1 tanta:2 tanto:3 tant’era:1 tardi,:1 te:1 tempo:3 tenere:1 tenni:1 terminava:1 terra:1 test’alta:1 ti:2 tolsi:1 tra:1 trarrotti:1 trattar:1 tremar:1 tremesse.:1 trovai,:1 tu:10 tua:1 tuo:2 tutta:1 tutte:2 tutt’i:1 udirai:1 umile:1 un:2 una:3 uscito:1 vagliami:1 valle:1 vederai:1 vedrai:1 veggia:1 vegna.:1 veltro:1 vene:1 venendomi:1 venire:1 venisse:1 venne:1 verace:1 vergine:1 vergognosa:1 verrà,:1 vestite:1 vi:1 via:3 via,:1 viaggio»,:1 vide,:1 vidi:2 villa,:1 virtute,:1 vissi:1 vista:1 vista,:1 vita:1 viva.:1 viver:1 voglia,:1 volge:1 volontieri:1 volse:1 volsi::1 volte:1 volto,:1 volume.:1 vorrai:1 vuol:1 vuo’:1 vòlto.:1 v’ho:1 v’intrai,:1 «A:1 «Miserere:1 «Non:1 «O:1 «Or:1 «Poeta,:1 «qual:1 «se:1 è:4 ’l:12 ’mpedisce:1 ’mpediva:1 ’n:3 ’ncontro,:1 ’nferno,:1 ’nvidia:1]
