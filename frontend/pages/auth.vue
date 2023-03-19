<template>
    <div>
        <AuthLogin v-if="states.login" />
        <AuthRegister v-if="states.register" />
    </div>
</template>

<script setup>
    const statesAuthentication = {
        login: true,
        register: false
    }
    const states = useState('statesAuthentication', () => statesAuthentication)

    const toggleLogin = () => {
        states.value.login = !states.value.login
        states.value.register = !states.value.register
    }

</script>