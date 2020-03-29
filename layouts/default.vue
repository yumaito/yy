<template>
  <v-app>
    <v-navigation-drawer
      v-model="drawer"
      clipped
      fixed
      app
    >
      <v-list-item>
        <v-list-item-content>
          <v-list-item-title class="title">
            メニュー
          </v-list-item-title>
        </v-list-item-content>
      </v-list-item>
      <v-divider />
      <v-list
        dense
        nav
      >
        <template v-for="item in items">
          <!-- サブメニューありの場合 -->
          <v-list-group
            v-if="item.lists"
            :key="item.title"
            :prepend-icon="item.icon"
            no-action
            :append-icon="item.lists ? undefined : ''"
          >
            <template v-slot:activator>
              <v-list-item-content>
                <v-list-item-title>{{ item.title }}</v-list-item-title>
              </v-list-item-content>
            </template>
            <v-list-item
              v-for="list in item.lists"
              :key="list.title"
              :to="list.to"
              nuxt
            >
              <v-list-item-title>{{ list.title }}</v-list-item-title>
              <v-list-item-icon>
                <v-icon>{{ list.icon }}</v-icon>
              </v-list-item-icon>
            </v-list-item>
          </v-list-group>
          <!-- サブメニューなしの場合 -->
          <v-list-item
            v-else
            :key="item.title"
            :to="item.to"
            nuxt
          >
            <v-list-item-action>
              <v-icon>{{ item.icon }}</v-icon>
            </v-list-item-action>
            <v-list-item-content>
              <v-list-item-title>
                {{ item.title }}
              </v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </template>
      </v-list>
    </v-navigation-drawer>
    <v-app-bar
      clipped-left
      flat
      dark
      fixed
      app
    >
      <v-app-bar-nav-icon @click.stop="drawer = !drawer" />
      <v-btn
        icon
        :to="{'name': 'index'}"
        nuxt
      >
        <v-icon>mdi-application</v-icon>
      </v-btn>
      <v-spacer />
      <v-btn
        text
      >
        プロフィール
      </v-btn>
      <v-btn
        text
        @click="logout"
      >
        <v-icon>mdi-logout-variant</v-icon>
        Logout
      </v-btn>
    </v-app-bar>
    <v-content>
      <v-container>
        <nuxt />
      </v-container>
    </v-content>
    <v-footer
      fixed
      app
    >
      <span>&copy; {{ new Date().getFullYear() }}</span>
    </v-footer>
  </v-app>
</template>

<script>
export default {
  middleware: 'auth',
  data () {
    return {
      drawer: null,
      items: [
        {
          icon: 'mdi-apps',
          title: 'ダッシュボード',
          to: { name: 'index' }
        },
        {
          icon: 'mdi-chart-bubble',
          title: 'Inspire',
          lists: [
            { title: 'Bingo', to: { name: 'inspire-bingo' } }
          ]
        }
      ]
    }
  },
  computed: {
    user () {
      return this.$auth.user
    }
  },
  methods: {
    logout () {
      this.$auth.logout()
    }
  }
}
</script>
