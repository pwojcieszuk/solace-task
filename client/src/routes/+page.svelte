<script lang="ts">
    import type { PageData } from "./$types";
    import AnimeComponent from "./AnimeComponent.svelte";

    export let data: PageData;
</script>

{#each data.recommended.data as recommendation}
    <div
        class="grid grid-cols-2 m-4 p-4 bg-yellow-300/90 max-w-[30rem] min-w-[20rem] slanted-corners"
    >
        {#each recommendation.entry as subRecommendation}
            <div class="p-2 text-xl font-bold">
                <AnimeComponent
                    title={subRecommendation.title}
                    mal_id={subRecommendation.mal_id}
                    image={subRecommendation.images.webp.image_url}
                />
            </div>
        {/each}

        <p class="col-span-2 p-2">
            {recommendation.content}
        </p>
    </div>
{/each}
