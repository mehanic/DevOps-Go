package main

import (
	"CAService/CAUtil"
	"CAService/CertGenUtil"
	"CAService/ClientUtil"
	"fmt"
)

func main() {
	// Create CA's private key and self-signed certificate
	caPrivateKey, caCert, err := CAUtil.CreateCACertificate()
	if err != nil {
		panic(err)
	}

	// Generate client's private key and CSR
	clientPrivateKey, clientCSRDER, err := ClientUtil.CreateClientCSR()
	if err != nil {
		panic(err)
	}

	// CA signs the client's CSR to issue a certificate
	err = CertGenUtil.SignCSR(caPrivateKey, caCert, clientCSRDER)
	if err != nil {
		panic(err)
	}

	fmt.Println(clientPrivateKey)

	println("CA and client certificates have been generated successfully.")
}


//&{{24279887109521942760210447900996697283120289596423639171340661287488433511183310270978071082766586051832888815957075119332168418453088586599933388091718670588121355524069973984712513130436570614878358944876392565385677360831390888004440661151447047972747552697148508819229558983528283446274718062995864750636757729343008149264673532189087798542817053603934211992199817224688732850234611107078635387192969242626653289065390792088779777677267440860435520280910552609303006798040158349955451781322355025408284862616110750913265701007711494063661174341204354570872641302357387851771493230793592772446781470309014023027263 65537} 17893250951610400101204574709283892844960003462594701080266281014741540835301150792189127930107581566617710057419393840326915639630243565552338110961042286831333219846827465454836885410894998055409844520189206221264284986701169066463806323943613371423589137436362340036180314017828845291792061476823386418759463654439708651655064222755864434988292614862735389711511570162844351745814676128514329090916699728344474408483401950600078087598645841029988798589888834136020267820454292768168610532894805045856356138837815111409004309512026994066544783027944433355671703639512290734298783513630740848686891748681259339492273 [175215780690975984711205471214714938434389068748340182463395750231285818443834768465477987436403291788696681134583052017342317553843187878041222666693151462067424769292458535701540164268235776494023799004367236399488674151040458470414956767305288282436046126053790288460895802473458284450642965474782606253101 138571349074680764372779184826425858146110221154457765193505126734347102902283452066228217773714692058880470020817396030745969234516793386088531470991086956744044239416609807473791767488364180582487237138310060614818158321955040932201341071716898728068537670346477230093085830887158733298369973356246959191963] {10982901675061863454134795241619985154774711909580564712446858140441615303832540837331333023921521013594854297585595582453304858495015270808754485478057680488471870123036142402946839110943628329606936025602951113555693324571985342576935813358715294631235446935761028197771639784563935372739693641308069433873 25544050966190370544693002912676057681357974606207245697891808231650019838602413665747640247863759323181332046347787684017304031649257379760067575583919336015148671077285549904494840212809980096999073986723894018457026880808380754412383866936415361304240407639284547915750948669419788165122108856322680531579 123242370853406915714516804248064019128439109104806275167277141679819288004417962293153552148407054363462834554315712657360936627901973170241386164556891866599956318282051758826869860601983711770028199173237949328454163268091490038655679016763138408778832791443568615971592297770437550607127483595450236613849 [] 0xc000080260 0xc000080280 0xc0000802a0}}
//CA and client certificates have been generated successfully.