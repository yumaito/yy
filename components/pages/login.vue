<template>
  <v-card
    :class="elevationClass"
  >
    <v-toolbar
      color="primary"
      dark
      flat
    >
      <v-toolbar-title>{{ title }}</v-toolbar-title>
    </v-toolbar>
    <v-card-text>
      <v-form
        ref="form"
        v-model="valid"
      >
        <v-text-field
          v-model="innerEmail"
          label="Email"
          name="email"
          prepend-icon="mdi-account"
          type="text"
          :rules="emailRules"
        />
        <v-text-field
          v-model="innerPassword"
          label="パスワード"
          name="password"
          prepend-icon="mdi-lock"
          type="password"
          :rules="passwordRules"
        />
      </v-form>
    </v-card-text>
    <v-card-actions>
      <v-spacer />
      <v-btn
        color="primary"
        @click="click"
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
  data () {
    return {
      valid: true,
      emailRules: [
        v => !!v || 'Email を入力してください',
        v => /.+@.+\..+/.test(v) || 'Email の形式が正しくありません'
      ],
      passwordRules: [
        v => !!v || 'password を入力してください'
      ]
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
  },
  methods: {
    click () {
      this.$refs.form.validate()
      if (this.valid) {
        this.onClick()
      }
    }
  }
}
</script>
