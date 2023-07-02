<script lang="ts">
    import type { LayoutData } from "../routes/$types";
    import AnimeComponent from "./AnimeComponent.svelte";
    import FavoritesToggleComponent from "./FavoritesToggleComponent.svelte";
    import { fly } from "svelte/transition";

    export let favorites: LayoutData["favorites"];

    let visible = true;

    const toggle = () => {
        visible = !visible;
    };
</script>

<button
    class="bg-gray-300 rounded-r-lg fixed w-[1rem] h-[4rem] left-0 top-[50%] translate-y-[-50%]"
    on:click={toggle}
    title="Toggle favorites"
>
    {#if visible}
        &larr;
    {:else}
        &rarr;
    {/if}
</button>

{#if visible}
    <div
        class:visible
        class="container"
        transition:fly={{ x: -300, duration: 1000 }}
    >
        {#each [...favorites] as [key, value]}
            <div class="relative">
                <FavoritesToggleComponent
                    title={value.title}
                    mal_id={key}
                    image={value.image}
                />
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
    </div>
{/if}
