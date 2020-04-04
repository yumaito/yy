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
          <login
            v-model="form"
            :elevation="24"
            :on-click="onClick"
          />
        </v-col>
      </v-row>
    </v-container>
  </v-content>
</template>

<script>
import Login from '@/components/pages/login'
export default {
  layout: 'login',
  components: {
    Login
  },
  data () {
    return {
      form: {
        email: '',
        password: ''
      }
    }
  },
  methods: {
    onClick () {
      this.$router.push({ name: 'index' })
    }
  }
}
</script>
