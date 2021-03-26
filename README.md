## Installation instructions ##

In order to run `fbz` you need the URL of the instance to talk to and the your user's token.  The first is basically the url you sign in via a web browser, while the second is available by going to the user options in the Fogbuz web UI.
With those in hand, you'll want to create file `~/.fbz.yml` with the format below:

```
url: https://<your_instance_name>.fogbugz.com
token: <your_user's_token>
```

If all went well, then running `fbz list` should provide a meaningful result.

## History ##

* v0.1.6 - Reparenting cases
* v0.1.5 - Show project and area for cases
* v0.1.4 - Accept/resolve cases
* v0.1.3 - Starting cases
* v0.1.2 - Assigning cases
* v0.1.1 - Estimating cases
* v0.1.0 - Initial releasev0.1.0 - Initial release
