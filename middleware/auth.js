export default function ({ store, redirect }) {
  if (!store.$auth.loggedIn) {
    redirect('/login')
  }
}
