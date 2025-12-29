Weather Server Project!

- This is a simple Golang HTTP server that serves the forecasted weather using the National Weather Service API.

Prerequisites

1. Clone the repository
*git bash terminal*
git clone https://github.com/Willintech/Jack-Henry.git
cd Jack-Henry

2. Run the server
command: go run main.go

3. Verify everything works
command: curl "http://localhost:8080/forecast?lat=38.8977&lon=-77.0365"

4. Response
{
  "forecast": "Partly cloudy, with a high near 75.",
  "temperature_type": "moderate"
}

