<template>
  <v-content>
    <v-container
      class="fill-height"
      fluid
    >
      <v-row
        align="center"
        justify="center"
      >
        <v-col
          cols="2"
          sm="8"
          md="4"
        >
          <v-card class="elevation-12">
            <v-toolbar
              dark
              flat
            >
              <v-toolbar-title>Login form</v-toolbar-title>
            </v-toolbar>
            <v-card-text>
              <v-form>
                <v-text-field
                  v-model="email"
                  label="Email"
                  name="email"
                  prepend-icon="mdi-account"
                  type="text"
                />
                <v-text-field
                  id="password"
                  v-model="pass"
                  label="Password"
                  name="password"
                  prepend-icon="mdi-lock"
                  type="password"
                />
              </v-form>
            </v-card-text>
            <v-card-actions>
              <v-spacer />
              <v-btn
                dark
                @click="login"
              >
                Login
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </v-content>
</template>

<script>
export default {
  layout: 'login',
  middleware ({ store, redirect }) {
    if (store.$auth.loggedIn) {
      redirect('/')
    }
  },
  data () {
    return {
      email: '',
      pass: ''
    }
  },
  methods: {
    async login () {
      try {
        await this.$auth.loginWith('local',
          {
            data: {
              email: this.email,
              password: this.password
            }
          })
      } catch (error) {
        console.log(error)
      }
    }
  }
}
</script>
