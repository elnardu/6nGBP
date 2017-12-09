
export default class Auth {
  constructor(vueHttp) {
    this.logged = false
    this.http = vueHttp
    this.token = ""
    this.admin = false

    if (localStorage.token) {
      this.http
        .post("/api/graphql", {
          query: `{ me { fullname } }`
        }, {
          headers: {
            "Authorization": "bearer " + localStorage.token
          }
        })
        .then(res => {
          if (res.body.data) {
            console.info('Token accepted')
            this.setToken(localStorage.token)
          } else {
            window.logged = false
          }
        })
    }
  }

  setToken(token) {
    localStorage.token = token
    this.token = token
    this.http.headers.common['Authorization'] = "bearer " + this.token
    this.logged = true
    window.logged = true

    this.checkIfAdmin()
  }

  logout() {
    localStorage.token = ""
    this.token = ""
    window.logged = false
    this.logged = false
    this.admin = false
    this.http.headers.common['Authorization'] = ""
  }

  checkIfAdmin() {
    this.http
      .post("/api/graphql", {
        query: `{ me { admin } }`
      })
      .then(res => {
        if (res.body.data) {
          this.admin = res.body.data.me.admin
        }
      })
  }
}
