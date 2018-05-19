package page

const Login = `
{{ define "screen" }}
<div id="login-page" class="row">

	<div class="col s12 z-depth-4 card-panel">
	
		<form class="login-form" action="/user/login" method="POST">

			<div class="row">
				<div class="input-field col s12 center">
					<img src="/images/login-logo.png" alt="" class="">
					<p class="center login-form-text">Teaching Mission</p>
				</div>
			</div>

			<div class="row margin">
				<div class="input-field col s12">
					<i class="mdi-communication-email prefix"></i>
					<input id="email" type="email" name="email" required>
					<label for="email" class="center-align">Email</label>
				</div>
			</div>

			<div class="row margin">
				<div class="input-field col s12">
					<i class="mdi-action-lock-outline prefix"></i>
					<input id="password" type="password" name="password" required>
					<label for="password">Password</label>
				</div>
			</div>

			<div class="row"> 
				<div class="input-field col s12 m12 l12  login-text">
					<input type="checkbox" id="remember" name="remember"/>
					<label for="remember">Remember me</label>
				</div>
			</div>

			<div class="row">
				<div class="input-field col s12">
					<!-- a href="/user/login" class="btn waves-effect waves-light col s12">Login</a -->
						<button class="btn waves-effect waves-light col s12" type="submit">Login</button>
				</div>
			</div>

			<div class="row">
				<div class="input-field col s6 m6 l6">
					<p class="margin medium-small"><a href="/user/register">Register Now!</a></p>
				</div>
				<div class="input-field col s6 m6 l6">
					<p class="margin right-align medium-small"><a href="/user/reset">Forgot password ?</a></p>
				</div>
			</div>

		</form>

	</div>

</div>
{{ end }}
`
