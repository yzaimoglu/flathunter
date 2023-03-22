<template>
    <div>
        <div>
            <div class="w-full overflow-hidden bg-white dark:bg-gray-900">

                <p class="px-8">
                    <span class="text-2xl font-bold text-gray-800 dark:text-white">Dashboard</span>
                </p>
                <div class="grid grid-cols-2 xl:grid-cols-4 gap-12 p-6">
                    <DashboardOption :option='{title: "URLs", description: "Change active URLs or add new URLs"}' />
                    <DashboardOption :option='{title: "URLs", description: "Change active URLs or add new URLs"}' />
                    <DashboardOption :option='{title: "URLs", description: "Change active URLs or add new URLs"}' />
                    <DashboardOption :option='{title: "URLs", description: "Change active URLs or add new URLs"}' />
                </div>
            </div>
        </div>
    </div>
</template>
