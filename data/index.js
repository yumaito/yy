const menues = [
  {
    icon: 'mdi-apps',
    title: 'ダッシュボード',
    to: { name: 'index' },
    hide: true
  },
  {
    icon: 'mdi-chart-bubble',
    title: 'Inspire',
    lists: [
      { title: 'Bingo', to: { name: 'inspire-bingo' } },
      { title: 'Game', to: { name: 'inspire-game' } }
    ]
  },
  {
    icon: 'mdi-account',
    title: 'ユーザー',
    to: { name: 'user' }
  }
]

export default {
  menues
}
