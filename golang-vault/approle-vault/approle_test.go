/*
Test Case:
  - confirm approle auth method was mounted
  - generate secret id
  - generate role id
  - generate token

Test Type: Unit Test
*/

package test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/hashicorp/vault/api"
	"github.com/hashicorp/vault/api/auth/approle"

	"github.com/palantir/stacktrace"
	"github.com/stretchr/testify/assert"
)

func TestApprole(t *testing.T) {

	// Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	// terraform.Options is a struct
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{

		// Set the path to the Terraform code that will be tested.
		TerraformDir: "../",
	})

	// Clean up resources with "terraform destroy" at the end of the test.
	defer terraform.Destroy(t, terraformOptions)

	// Run "terraform init" and "terraform apply". Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	/*
	 construct approle path: in the data structure, all auth methods in sys/auths have a "/". for example: mount path is "approle/"
	*/
	expectedApprolePath := "app-demo"
	approlePath := expectedApprolePath + "/"

	// VAULT CLIENT
	config := api.DefaultConfig()        // returns default configuration of the vault client
	client, err := api.NewClient(config) // create and returns a vault client
	// error handling
	if err != nil {

		err = stacktrace.Propagate(err, "initialization error. Could not setup client with default config")
		fmt.Println(err)
		os.Exit(1)

	}

	/*
		Exexute Test Cases
	*/

	// store state of approle mount
	approleExist, _ := isApproleMounted(client, approlePath)
	// assert approle was enabled
	assert.True(t, approleExist, "approle doesn't exist")

	// generate secret-id
	roleName := "myrole"
	secretID := generateSecretID(client, approlePath, roleName)

	// generate role-id
	roleID := generateRoleID(client, approlePath, roleName)

	if roleID == "" || secretID == "" {

		t.Error("either role-id, secret-id or both was not generated")

	} //else {
	// 	// generate token
	/*token :=*/
	generateToken(t, client, roleID, secretID)
	// 	if token == "" {
	// 		t.Error("no token was generated")
	// 	}

	// }

}

// function returns data about the approle and the boolean state of the approle mount
func isApproleMounted(c *api.Client, approlePath string) (bool, *api.MountOutput) {

	// mountData contains information about the mounted auth method
	var mountData *api.MountOutput

	// returns all enabled auth methods
	auths, err := c.Sys().ListAuth()
	if err != nil {

		err = stacktrace.Propagate(err, "Could not find any auth method")
		fmt.Println(err)
		os.Exit(1)

	}

	// traverse the returned auth methods
	for auth, mountData := range auths {

		if auth == approlePath {

			return true, mountData

		}

	}

	return false, mountData

}

// generate secret-id
func generateSecretID(c *api.Client, approlePath string, roleName string) string {

	var secretID string // secret-id

	// approle auth backend path for secret id
	path := "auth/" + approlePath + "role/" + roleName + "/secret-id"

	// write to backend to generate secret-id
	secret, err := c.Logical().Write(path, map[string]interface{}{})
	if err != nil {

		err = stacktrace.Propagate(err, "no secret")
		fmt.Println(err)
		os.Exit(1)

	}

	// extract secret-id
	for key, id := range secret.Data {

		if key == "secret_id" {

			// type assertion provides access to an interface's value
			secretID = id.(string)

		}

	}

	return secretID

}

// generate role-id
func generateRoleID(c *api.Client, approlePath string, roleName string) string {

	var roleID string // secret-id

	// approle auth backend path for secret id
	path := "auth/" + approlePath + "role/" + roleName + "/role-id"

	// write to backend to generate secret-id
	role, err := c.Logical().Read(path)
	if err != nil {

		err = stacktrace.Propagate(err, "no role-id")
		fmt.Println(err)
		os.Exit(1)

	}

	// extract secret-id
	for key, id := range role.Data {

		if key == "role_id" {

			// type assertion provides access to an interface's value
			roleID = id.(string)

		}

	}

	return roleID

}

// generate token
func generateToken(t *testing.T, c *api.Client, roleID string, secretID string) {

	// initialize SecretID struct
	secret := &approle.SecretID{
		FromString: secretID,
	}

	// initializes new approle auth method interface
	approleAuth, err := approle.NewAppRoleAuth(roleID, secret)
	if err != nil {

		err = stacktrace.Propagate(err, "could not create approle auth method interface")
		fmt.Println(err)
		os.Exit(1)

	}

	// create context
	ctx := context.Background()

	// generate token for vault client
	login, _ := approleAuth.Login(ctx, c)
	// login.Auth.ClientToken

	if login == nil {
		t.Errorf("wnat %#v", login)
	}

}

