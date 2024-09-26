import moment from "moment";

export default defineNuxtPlugin((nuxtApp) => {
    return {
        provide: {
            moment:(d) => moment().from(d)
        }
    }
})