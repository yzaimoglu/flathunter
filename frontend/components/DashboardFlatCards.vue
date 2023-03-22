<template>
    <div>
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 justify-center gap-12">
            <DashboardFlatCard v-for="flat in flats" v-bind:key="flat.id" :flat="flat" class="flex justify-center"/>
        </div>
    </div>
</template>

<script setup>
    const flats = ref(
        [
            {id: 1, name: "test", price: 500},
            {id: 2, name: "test 2", price: 200},
            {id: 3, name: "test 3", price: 1000},
            {id: 4, name: "test 4", price: 100},
            {id: 5, name: "test 5", price: 100},
            {id: 6, name: "test 6", price: 100},
            {id: 7, name: "test 7", price: 100},
            {id: 8, name: "test 8", price: 100},
            {id: 9, name: "test 9", price: 100},
            {id: 10, name: "test 10", price: 100},
        ]
    )
</script>
