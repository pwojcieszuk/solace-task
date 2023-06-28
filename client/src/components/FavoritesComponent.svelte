<script lang="ts">
    import type { LayoutData } from "../routes/$types";
    import AnimeComponent from "./AnimeComponent.svelte";
    export let favorites: LayoutData["favorites"];
</script>

{#each [...favorites] as [key, value]}
    <div class="relative">
        <div class="star" />
        <div
            class="m-4 p-4 padding-4 rounded-full h-[200px] w-[200px] bg-yellow-300/90 bg-circle flex justify-center items-center"
        >
            <AnimeComponent
                title={value.title}
                mal_id={key}
                image={value.image}
                height="150px"
                width="150px"
            />
        </div>
    </div>
{/each}
