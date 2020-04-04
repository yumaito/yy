<template>
  <v-card
    :class="elevationClass"
  >
    <v-toolbar
      dark
      flat
    >
      <v-toolbar-title>{{ title }}</v-toolbar-title>
    </v-toolbar>
    <v-card-text>
      <v-form>
        <v-text-field
          v-model="innerEmail"
          label="Email"
          name="email"
          prepend-icon="mdi-account"
          type="text"
        />
        <v-text-field
          v-model="innerPassword"
          label="パスワード"
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
        @click="onClick"
      >
        {{ btnText }}
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  name: 'Login',
  props: {
    value: {
      type: Object,
      required: true,
      default () {
        return {
          email: '',
          password: ''
        }
      }
    },
    elevation: {
      type: Number,
      required: false,
      default: 12
    },
    onClick: {
      type: Function,
      required: true
    },
    title: {
      type: String,
      required: false,
      default: 'ログイン'
    },
    btnText: {
      type: String,
      required: false,
      default: 'ログイン'
    }
  },
  computed: {
    elevationClass () {
      return [
        `elevation-${this.elevation}`
      ]
    },
    innerEmail: {
      get () {
        return this.value.email
      },
      set (value) {
        const f = {
          email: value,
          password: this.value.password
        }
        this.$emit('input', f)
      }
    },
    innerPassword: {
      get () {
        return this.value.password
      },
      set (value) {
        const f = {
          email: this.value.email,
          password: value
        }
        this.$emit('input', f)
      }
    }
  }
}
</script>
