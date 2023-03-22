<template>
    <div>
        <div class="cursor-pointer justify-center p-6 rounded-lg shadow bg-white dark:bg-gray-700">
            <span class="inline-block p-3 text-blue-500 rounded-full bg-blue-100/80 dark:bg-gray-800">
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
                    stroke="currentColor" class="w-6 h-6">
                    <path stroke-linecap="round" stroke-linejoin="round"
                        d="M21.75 6.75v10.5a2.25 2.25 0 01-2.25 2.25h-15a2.25 2.25 0 01-2.25-2.25V6.75m19.5 0A2.25 2.25 0 0019.5 4.5h-15a2.25 2.25 0 00-2.25 2.25m19.5 0v.243a2.25 2.25 0 01-1.07 1.916l-7.5 4.615a2.25 2.25 0 01-2.36 0L3.32 8.91a2.25 2.25 0 01-1.07-1.916V6.75" />
                </svg>
            </span>

            <h2 class="mt-4 text-lg font-medium text-gray-800 dark:text-white">{{ props.option.title }}</h2>
            <p class="mt-2 mb-2 text-gray-500 dark:text-gray-400">{{ props.option.description }}</p>
        </div>
        <!-- <div>
            <button @click="testToast">Test Toast</button>
        </div>
        <div>
            <button @click="testDialog">Test Dialog</button>
        </div> -->
    </div>
</template>

<script setup>
    import { Toast } from '/utilities/toast.js'
    import { Dialog } from '/utilities/dialog.js'

    const props = defineProps({
        option: {
            type: Object,
            required: true
        },
    })

    const testToast = async () => {
        await Toast('tr', 'Success', 'success')
        await Toast('br', 'Question', 'question')
        await Toast('bl', 'Error', 'error')
        await Toast('tr', 'Warning', 'warning')
        await Toast('br', 'Info', 'info')
    }

    const testDialog = async () => {
        const result = await Dialog(['1', '2', '3'])
        console.log(result)
    }
</script>
