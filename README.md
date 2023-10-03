# DDRest
### Go / Htmx
Another day, another Htmx exercise. Today I tried to build a simple form that filters out rows on a real estate listing webpage. This is the second time I've tried to use json with Htmx, but this time I did something different. Instead of just using the json-enc extension for Htmx, I actually built on top of it, adding a data parser to convert digits to numbers and booleans. This made my life a little easier on the backend, since I didn't have to implement a custom json parser in Go.

### Json
On POST the form then sends a json with only the fields that have values in it, the backend parses it and queries the sqlite database for the rows that match, then it responds with html and htmx takes care to append it to the DOM.

### Backend
This exercise is all about improving my Go skills by designing and implementing a small MVC using only Go's standard library. That's why I'm sticking with JSON and htmx, because they're just simpler to implement in Go.

### Beyond Htmx
One thing I really wanted to try with this exercise was to save the contents of the form so that it would stay filtered on page reload. In Django, I would just use request.session, but with Go, I would need a third-party package just for sessions. So, I decided to save the contents of the form in the browser's localStorage instead.
I couldn't figure out how to run a function on page load with just htmx, and I didn't want to just use a `document.addEventListener`, so the solution that I found was to use alpine.js, which seems to work perfectly with Htmx. I'm probably going to keep using it on other projects and learn more about it.

### How to run

All you'll need is go 1.21, cd into the folder and call `go run .`

### Conclusion

I've been wanting to try this for a few weeks now, and the result, although simple, made me really happy. I could expand it by paginating the table, but I'll save that for next time.
