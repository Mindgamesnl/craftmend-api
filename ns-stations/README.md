# NS API
A simple REST api that relays real time train departure details from any given station. It does this by faking a api key, making a few requests, processing the data and then presenting it to you. You have to provide a station id, which you have to figure out yourself.

You should do a GET request to `https://api.craftmend.com/ns/<STATION ID>`, with a example for Hilversum Statino being `https://api.craftmend.com/ns/hvs`

## Notes
 - I don't own this data
 - I'm not responsible for this data
 - The request INCLUDES a cors wildcard, meaning that you can use this api in your own systems.
