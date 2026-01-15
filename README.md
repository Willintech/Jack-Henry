Weather Server Project!

This is a Golang HTTP server that provides forecasted weather information using the National Weather Service (NWS) API. https://www.weather.gov/documentation/services-web-api 
The server should expose a single endpoint that accepts latitude and longitude coordinates and returns:
- The short forecast for today (e.g., "Partly Cloudy").
- A characterization of the temperature: "hot", "moderate", or "cold".

Prerequisites
*git bash terminal*
*Make sure go for windows is fully installed, go version go1.25.5*
- https://go.dev/dl/      (click download)
- export PATH="/c/Program Files/Go/bin:$PATH     (in bash terminal run this command after the download is complete)
- command:  go version          (run this command to check go is installed)

1. Clone the repository
- git clone https://github.com/Willintech/Jack-Henry.git
- cd Jack-Henry   (folder & project)

2. Run the server
- command: go run main.go
- a windows security prompt will pop up: select "allow"
- you will see: Starting server on :8080

4. Test to Verify everything works 
run this command in a new POWERSHELL terminal: curl.exe "http://localhost:8080/forecast?lat=38.8977&lon=-77.0365"
                                  *OR*
On your internet browser: http://localhost:8080/forecast?lat=38.8977&lon=-77.0365

5. Successful Response *Example*
{
  "forecast": "Partly cloudy, with a high near 75.",
  "temperature_type": "moderate"
}


Fascinating & very interesting project!


to exit/quit server run this command in the bash terminal: ctrl + c         (control button and c key)
