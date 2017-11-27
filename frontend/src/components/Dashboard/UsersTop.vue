<template>
    <div class="box" :class="{'is-loading': loading}">
        <table class="table is-striped is-fullwidth">
            <thead>
                <th>Name</th>
                <th>Points</th>
            </thead>
            <tbody>
                <tr v-for="u in users" :key="u.id">
                    <td>{{ u.fullname }}</td>
                    <td>{{ u.points }}</td>
                </tr>
            </tbody>
        </table>
    </div>
</template>

<script>
export default {
    data() {
       return {
           loading: true,
           users: []
       }
    },
    mounted() {
        this.load()
    },
    methods: {
        load() {
            this.$http.post('/api/graphql', {
                query: `
                    {
                        users {
                            fullname
                            points
                            id
                        }
                    }
                `
            }).then(res => {
                console.log(res)
                if (res.body.errors) {
                    this.$emit('error', res.body.errors)
                    console.error(res.body.errors)
                    return
                }
                this.users = res.body.data.users
                this.loading = false
            })
        }
    }
}
</script>

<style scoped>
.box {
    margin-left: 20px;
    margin-right: 20px;
}
</style>
