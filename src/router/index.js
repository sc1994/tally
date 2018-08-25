import Vue from "vue"
import Router from "vue-router"
import Sign from "@/pages/Sign"
import Home from "@/pages/Home"
import Me from "@/pages/Me"
import Tally from "@/pages/Tally"
import AddPartner from "@/pages/AddPartner"
import Message from "@/pages/Message"

Vue.use(Router)

export default new Router({
  routes: [{
      path: "/sign",
      name: "Sign",
      component: Sign
    },
    {
      path: "/",
      name: "Home",
      component: Home
    },
    {
      path: "/tally",
      name: "Tally",
      component: Tally
    },
    {
      path: "/me",
      name: "Me",
      component: Me
    },
    {
      path: "/addpartner",
      Name: "AddPartner",
      component: AddPartner
    },
    {
      path: "/message",
      Name: "Message",
      component: Message
    }
  ]
})
