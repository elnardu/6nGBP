<template>
    <div class="box" :class="{'is-loading': loading}">
        <table class="table is-striped is-fullwidth">
            <thead>
                <th>User</th>
                <th>Category</th>
                <th>Added By</th>
                <th>Points</th>
            </thead>
            <tbody>
                <tr v-for="a in actions" :key="a.id">
                    <td>{{ a.user.fullname }}</td>
                    <td>{{ a.category.title }}</td>
                    <td>{{ a.addedBy.fullname }}</td>
                    <td>{{ a.points }}</td>
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
           actions: []
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
                        actions {
                            points
                            id
                            user {
                                fullname
                            }
                            addedBy {
                                fullname
                            }
                            category {
                                title
                            }
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
                this.actions = res.body.data.actions
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
