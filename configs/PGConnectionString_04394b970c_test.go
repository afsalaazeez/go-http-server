/*
Test generated by RoostGPT for test bitci using AI Type Open AI and AI Model gpt-4

1. Test scenario: Validate the connection string when all the required environment variables (DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME) are set. The expected output should be a string in the format "postgres://user:password@host:port/database?sslmode=disable".

2. Test scenario: Validate the connection string when none of the environment variables are set. The function should use the default values for DB_USER, DB_PORT, and DB_NAME ("postgres", "5432", "postgres" respectively). An error should be returned due to missing required environment variables (DB_PASSWORD, DB_HOST).

3. Test scenario: Validate the connection string when some of the required environment variables are not set. For example, if DB_PASSWORD and DB_HOST are not set, an error should be returned.

4. Test scenario: Validate the function when non-default values are set for the environment variables. The function should return a connection string with these values.

5. Test scenario: Validate the connection string when the DB_PORT is set to a non-numeric value. An error should be returned since the DB_PORT should be an integer.

6. Test scenario: Validate the function with edge case values for the environment variables. For example, empty strings or very long strings.

7. Test scenario: Validate the function when the environment variables contain special characters. The function should handle these correctly and return a valid connection string.

8. Test scenario: Validate the function when the DB_USER, DB_PASSWORD, DB_HOST, and DB_NAME are set to null. An error should be returned.

9. Test scenario: Validate the connection string when the DB_PORT is set to a negative number. An error should be returned since the port number should be a positive integer.

10. Test scenario: Validate the function when the DB_PORT is set to zero. An error should be returned since the port number should be a positive integer.
*/
package configs

import (
	"fmt"
	"os"
	"testing"
)

func TestPGConnectionString_04394b970c(t *testing.T) {
	// Test cases
	testCases := []struct {
		name           string
		dbConfig       *DBConfigs
		expectedOutput string
		expectedError  error
		envVariables   map[string]string
	}{
		{
			name: "Test when all environment variables are set",
			dbConfig: &DBConfigs{
				User:     "testuser",
				Password: "testpass",
				Host:     "testhost",
				Port:     5432,
				Database: "testdb",
			},
			expectedOutput: "postgres://testuser:testpass@testhost:5432/testdb?sslmode=disable",
			envVariables: map[string]string{
				"DB_USER":     "testuser",
				"DB_PASSWORD": "testpass",
				"DB_HOST":     "testhost",
				"DB_PORT":     "5432",
				"DB_NAME":     "testdb",
			},
		},
		{
			name: "Test when no environment variables are set",
			dbConfig: &DBConfigs{
				User:     "postgres",
				Port:     5432,
				Database: "postgres",
			},
			expectedError: fmt.Errorf("required environment variables are missing"),
			envVariables:  map[string]string{},
		},
		// Add more test cases here
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			// Set environment variables based on test case
			for key, val := range test.envVariables {
				os.Setenv(key, val)
			}

			// Call function
			result := test.dbConfig.PGConnectionString()

			// Compare result with expected
			if result != test.expectedOutput {
				t.Errorf("expected '%s', got '%s'", test.expectedOutput, result)
			}

			// Reset environment variables
			for key := range test.envVariables {
				os.Unsetenv(key)
			}
		})
	}
}
