<script lang="ts">
    import "../app.css";
    import type { LayoutData } from "./$types";
    import AnimeComponent from "./AnimeComponent.svelte";
    export let data: LayoutData;
</script>

<div class="bg-hive-pattern bg-cover h-screen overflow-hidden">
    <div class="overflow-y-auto h-screen">
        <div class="flex flex-wrap justify-center"><slot /></div>

        <div>
            {#each [...data.favorites] as [key, value]}
                <AnimeComponent
                    title={value.title}
                    mal_id={key}
                    image={value.image}
                />
            {/each}
        </div>
    </div>
</div>
