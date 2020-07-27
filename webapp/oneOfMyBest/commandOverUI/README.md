# Command Over UI

> ./programName

Makesure, you have `commands.html` and `.htaccess`

### Listen at port 80

Over here I have build a code, that will help you define your commands that you want to run.

Note:
commands that are allowed to the user, who is running this code only will get executed, as as per my initial test, you can't use `;` to run multiple commands.

## Features:
- Password protected using .htaccess file. 
- Commands can be composed at the `commands.html`
- The file need to be at the same place.


### Example of `commands.html`

```
<!DOCTYPE html>
<html>
<body>

<h1>The button value attribute</h1>

<!--
Note: Get is not safe, use post method.
Either order is good.
<form action="/dnsCheck" method="post">
<form method="post" action="/dnsCheck">
-->

<form method="post" action="/">
<br>
<p>
Commands to run on the hosted server:
<br>
<button name="command" type="submit" value="ls">ls</button>
<button name="command" type="submit" value="pwd">pwd</button>
<button name="command" type="submit" value="uptime">uptime</button>
<button name="command" type="submit" value="ps -ef">ps -ef</button>
</p>
</form>

</body>
</html>
```


Over here the example commands are:

```
<button name="command" type="submit" value="ls">ls</button>
<button name="command" type="submit" value="pwd">pwd</button>
<button name="command" type="submit" value="uptime">uptime</button>
<button name="command" type="submit" value="ps -ef">ps -ef</button>
```

If you want to add or remove the command: just add the same. You can also create different sessions:

```
<button name="command" type="submit" value="ls">ls</button>
<button name="command" type="submit" value="pwd">pwd</button>
<button name="command" type="submit" value="uptime">uptime</button>
<button name="command" type="submit" value="ps -ef">ps -ef</button>

<button name="command" type="submit" value="foo">foo</button>
<button name="command" type="submit" value="your_command_or_script">your_command_script</button>
```

Note:

Make sure, your script are in the executable path or you might like to provide the same in the correct place/location.


##  Password protected using .htaccess file.

This program expect the standard .htaccess file, which you can create as:

When creating for the first time.
> htpasswd -c .htpasswd amitmund

If you want to add a new user
> htpasswd .htpasswd userFoo

To update the password
> htpasswd .htpasswd userFoo
And give the new password.

if you want, feel free to learn more details on .htpasswd


## Todo:
Make it SSL based and with JWT.


### PS
Its not perfect, but help you in your regular works.


### ngrok

If you want:
>ngrok http 80