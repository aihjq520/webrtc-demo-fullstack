import Client from "../components/client.vue"
import Remote from "../components/remote.vue"
import Login from "../views/login/index.vue"

const routes = [
    {
        path: '/',
        name: 'login',
        component: Login
    },
    {
        path: '/offer',
        name: 'offer',
        component: Client
    },
    {
        path: '/answer',
        name: 'answer',
        component: Remote
    }
]

export default routes