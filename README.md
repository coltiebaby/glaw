# g-LAW - Go(Lang) League API Wrapper

Wanted to make a simple wrapper to avoid making those weird urls RIOT wants.

*Note:* I have no idea what I'm doing in go; please take all this with a grain
of salt. If you have a suggestion please feel free to open a PR and explain
why you're making a change. I'll greatly appreciate it.

## ROADMAP

1. Add in all the api callings

   - ~~champions~~
   - ~~champion-masteries~~
   - ~~league~~
   - ~~masteries~~
   - match
   - runes
   - status
   - static-data
     - champions
     - items
     - language-strings
     - languages
     - maps
     - masteries
     - profile-icons
     - realms
     - runes
     - summoner-spells
     - versions
   - spectator
   - ~~summoner~~
   - tournaments-stub
   - tournament

2. Errors raise a 404 page with a detailed explanation

3. Caching to lower the api requests.

4. Small docker image that runs the latest release.

5. Your own authenication with rate limiting
