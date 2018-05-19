package page

const Register3 = `
{{ define "screen" }}
<div id="login-page" class="row">

	<div class="col s12 z-depth-4 card-panel">

		<form method="POST" class="login-form" name="register" action="/user/register">
			<input id="step" name="step" type="hidden" ref="step" value="3">

			<div class="row">
				<div class="input-field col s12 center">
					<h4>Register (cont.)</h4>
				</div>
			</div>

			<div class="row margin">
				<div class="input-field col s12">
					<input id="profession" name="profession" type="text" v-model="profession" ref="profession" required>
					<label for="profession" class="center-align">Profession</label>
				</div>
			</div>

			<div class="row margin">
                <div class="input-field col s12">
                  <textarea id="introduction" name="introduction" class="materialize-textarea" v-model="introduction" ref="introduction" required></textarea>
                  <label for="introduction" class="">Introduce Yourself</label>

                </div>
			</div>

			<div class="row margin">
                <div class="input-field col s12">
                  <input type="checkbox" name="ishuman" class="filled-in" v-model="isHuman" id="isHuman" ref="isHuman" required/>
                  <label for="isHuman">I am a Human Being</label>
                </div>
			</div>

			<br>

			<div class="row">
				<div class="input-field col s12">
					<button class="btn waves-effect waves-light col s12" type="submit">Register</button>
				</div>
				<div class="input-field col s12">
					<p class="margin center medium-small sign-up">Already have an account? <a href="/user/login">Login</a></p>
				</div>
			</div>
			
		</form>

	</div>
</div>
{{ end }}
`
