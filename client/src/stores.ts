import { writable } from "svelte/store";

function createFavorites() {
    const { subscribe, set, update } = writable(
        new Map<number, { title: string; image: string }>(),
    );

    const fv = [
        {
            id: 51252,
            title: "Spy Kyoushitsu",
            image: "https://cdn.myanimelist.net/images/anime/1491/132864.webp",
        },

        {
            id: 51815,
            title: "Kubo-san wa Mob wo Yurusanai",
            image: "https://cdn.myanimelist.net/images/anime/1818/132330.webp",
        },

        {
            id: 24833,
            title: "Ansatsu Kyoushitsu",
            image: "https://cdn.myanimelist.net/images/anime/5/75639.webp",
        },
    ];

    const fetchFavorites = async () => {
        const data = await Promise.resolve(fv);

        set(
            new Map(
                data.map((f: { id: number; title: string; image: string }) => [
                    f.id,
                    { title: f.title, image: f.image },
                ]),
            ),
        );
    };

    fetchFavorites();

    return {
        subscribe,
        add: (mal_id: number, title: string, image: string) =>
            update((n) => n.set(mal_id, { title: title, image: image })),
        remove: (mal_id: number) =>
            update((n) => {
                n.delete(mal_id);
                return n;
            }),
    };
}

export const favorites = createFavorites();
