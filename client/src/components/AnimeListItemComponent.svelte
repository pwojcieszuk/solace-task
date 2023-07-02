<script lang="ts">
    import AnimeComponent from "./AnimeComponent.svelte";
    import FavoritesToggleComponent from "./FavoritesToggleComponent.svelte";
    export let recommendation: {
        content: string;
        entry: {
            mal_id: number;
            title: string;
            images: {
                webp: {
                    image_url: string;
                };
            };
        }[];
    };
</script>

<div
    class="grid grid-cols-2 m-4 p-4 bg-yellow-300/90 max-w-[30rem] min-w-[20rem] slanted-corners"
>
    {#each recommendation.entry as subRecommendation}
        <div class="p-2 text-xl font-bold relative">
            <FavoritesToggleComponent
                title={subRecommendation.title}
                mal_id={subRecommendation.mal_id}
                image={subRecommendation.images.webp.image_url}
            />
            <AnimeComponent
                title={subRecommendation.title}
                mal_id={subRecommendation.mal_id}
                image={subRecommendation.images.webp.image_url}
            />

            <h1>{subRecommendation.title}</h1>
        </div>
    {/each}

    <p class="col-span-2 p-2">
        {recommendation.content}
    </p>
</div>
