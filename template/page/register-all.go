package page

const Register = `
{{ define "screen" }}
<div id="login-page" class="row">

	<div class="col s12 z-depth-4 card-panel">

		<form method="POST" class="login-form" name="register" action="/user/register">

			<div class="row">
				<div class="input-field col s12 center">
					<h4>Register {</h4>
					<p class="margin medium-small red-text">
					</p>
				</div>
			</div>

			<div id="screen1" :class="showScreen1">
				<div class="row margin">
					<div class="input-field col s6">
						<i class="mdi-social-person-outline prefix"></i>
						<input id="fname" name="fname" type="text" v-model="fname" ref="fname" required>
						<label for="fname" class="center-align">First Name</label>
					</div>
					<div class="input-field col s6">
						<input id="lname" name="lname" type="text" v-model="lname" ref="lname" required>
						<label for="lname" class="center-align">Last Name</label>
					</div>
				</div>

				<div class="row margin">
					<div class="input-field col s12">
						<i class="mdi-communication-email prefix"></i>
						<input id="email" name="email" type="email" v-model="email" ref="email" required>
						<label for="email" class="center-align">Email</label>
					</div>
				</div>

				<div class="row margin">
					<div class="input-field col s12">
						<i class="mdi-action-lock-outline prefix"></i>
						<input id="password" name="password" type="password" v-model="password" ref="password" required>
						<label for="password">Password</label>
					</div>
				</div>

				<div class="row margin">
					<div class="input-field col s12">
						<i class="mdi-action-lock-outline prefix"></i>
						<input id="password-again" type="password" v-model="passwordagain" ref="passwordagain" required>
						<label for="password-again">Password again</label>
					</div>
				</div>

				<div class="row">
					<div class="input-field col s12">
						<a @click="submit1()" class="btn waves-effect waves-light col s12">Next</a>
					</div>
					<div class="input-field col s12">
						<p class="margin center medium-small sign-up">Already have an account? <a href="/user/login">Login</a></p>
					</div>
				</div>
			</div>


			<div id="screen2" :class="showScreen2">
				<div class="row margin">
					<div class="input-field col s12">
						<i class="mdi-social-cake prefix"></i>
						<input id="dob" name="dob" type="date" class="datepicker" ref="dob">
						<label for="dob" class="center-align">Date of Birth</label>
					</div>
				</div>

				<div class="row margin">
					<div class="input-field col s12">
						<input id="address" name="address" type="text" v-model="address" ref="address" required>
						<label for="address" class="center-align">Address</label>
					</div>
				</div>

				<div class="row margin">
					<div class="input-field col s12">
						<input id="city" name="city" type="text" v-model="city" ref="city" required>
						<label for="city" class="center-align">City</label>
					</div>
				</div>

				<div class="row margin">
					<div class="input-field col s6">
						<input id="state" name="state" type="text" v-model="state" ref="state" required>
						<label for="state" class="center-align">State</label>
					</div>
					<div class="input-field col s6">
						<input id="country" name="country" type="text" v-model="country" ref="country" required>
						<label for="country" class="center-align">Country</label>
					</div>
				</div>
				
				<div class="row">
					<div class="input-field col s4">
						<a @click="submit0()" class="btn waves-effect waves-light col s12">Back</a>
					</div>
					<div class="input-field col s8">
						<a @click="submit2()" class="btn waves-effect waves-light col s12">Next</a>
					</div>
					<div class="input-field col s12">
						<p class="margin center medium-small sign-up">Already have an account? <a href="/user/login">Login</a></p>
					</div>
				</div>
			</div>


			<div id="screen3" :class="showScreen3">
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
					<div class="input-field col s4">
						<a @click="submit1()" class="btn waves-effect waves-light col s12">Back</a>
					</div>
					<div class="input-field col s8">
						<!--a @click="submit3()" class="btn waves-effect waves-light col s12">Register</a -->
                        <!-- button class="btn waves-effect waves-light col s12" type="submit">Register</button -->
						<button @click="submit3()" class="btn waves-effect waves-light col s12" type="submit">Register</button>

					</div>
					<div class="input-field col s12">
						<p class="margin center medium-small sign-up">Already have an account? <a href="/user/login">Login</a></p>
					</div>
				</div>
			</div>

			
		</form>

	</div>
</div>
{{ end }}

{{ define "javascript" }}
<script>
	var app = new Vue({
		el: '#login-page',
		data: {
			title: "",
			error: '',
			showScreen1: 'show',
			showScreen2: 'hide',
			showScreen3: 'hide',

			// Screen 1
			fname: '',
			fnameClass: '',
			lname: '',
			email: '',
			password: '',
			passwordagain: '',

			// Screen 2
			address: '',
			city: '',
			state: '',
			country: '',

			// Screen 3
			profession: '',
			introduction: '',
			isHuman: false

		},
		methods: {
			submit0: function() {
				this.title = "Register"
				this.showScreen1 = 'show'
				this.showScreen2 = 'hide'
				this.showScreen3 = 'hide'
			},
			submit1: function() {
				if (this.fname.length == 0) { this.$refs.fname.focus(); return}
				if (this.lname.length == 0) { this.$refs.lname.focus(); return}
				if (this.email.length == 0) { this.$refs.email.focus(); return}
				if (this.password.length < 8) {
					this.error = 'Password should be atleast 8 characters'
					this.$refs.password.focus();
					return
				}
				if (this.password != this.passwordagain) {
					this.error = 'Passwords do not match'
					this.$refs.passwordagain.focus(); 
					return
				}

				this.error = ''
				this.title = "(cont.)"
				this.showScreen1 = 'hide'
				this.showScreen2 = 'show'
				this.showScreen3 = 'hide'
			},
			submit2: function() {
				if (this.address.length == 0) { this.$refs.address.focus(); return}
				if (this.city.length == 0) { this.$refs.city.focus(); return}
				if (this.state.length == 0) { this.$refs.state.focus(); return}
				if (this.country.length == 0) { this.$refs.country.focus(); return}

				this.error = ''
				this.title = "(cont.)"
				this.showScreen1 = 'hide'
				this.showScreen2 = 'hide'
				this.showScreen3 = 'show'
			},
			submit3: function() {
				if (this.profession.length == 0) { this.$refs.profession.focus(); return}
				if (this.introduction.length == 0) { this.$refs.introduction.focus(); return}
				if (!this.isHuman) {
					this.error = 'Please accept "I am a Human Being"'
					this.$refs.introduction.focus(); 
					return
				}
				this.error = ''
			}
		}
	})
</script>
{{ end }}
`
