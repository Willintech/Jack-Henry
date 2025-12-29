Jack Henry Weather Server Project!

This is a Golang HTTP server that provides forecasted weather information using the National Weather Service (NWS) API.  
The server exposes a single endpoint that accepts latitude and longitude coordinates and returns:
- The short forecast for today (e.g., "Partly Cloudy").
- A characterization of the temperature: "hot", "moderate", or "cold".

Prerequisites
*git bash terminal*
*Make sure go is fully installed, go version go1.25.5*

1. Clone the repository
git clone https://github.com/Willintech/Jack-Henry.git
cd Jack-Henry

2. Run the server
command: go run main.go
you will see: Starting server on :8080

3. Verify everything works
command: curl.exe "http://localhost:8080/forecast?lat=38.8977&lon=-77.0365"
                            OR
In your browser: http://localhost:8080/forecast?lat=38.8977&lon=-77.0365

4. Successful Response *Example*
{
  "forecast": "Partly cloudy, with a high near 75.",
  "temperature_type": "moderate"
}

Fascinating & very interesting project. 
