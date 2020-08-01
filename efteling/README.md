# Efteling API
A simple rest api that forges efteling api requests and returns the current state of all rides, shows, and more. The data updates every 5 minutes and gets cached on my side.

Just make a simple GET request to `https://api.craftmend.com/themeparks/efteling` and parse the resulting JSON, and that's it.

## Notes
 - I don't own this data
 - I'm not responsible for this data
 - The request INCLUDES a cors wildcard, meaning that you can use this api in your own systems.
