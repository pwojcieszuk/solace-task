import { PUBLIC_SERVER_API_URL } from "$env/static/public";

const apiUrl = PUBLIC_SERVER_API_URL;

export const load = async () => {
    const res = await fetch(`${apiUrl}/favorites`);
    const favorites = new Map(
        (await res.json()).map(
            (f: { ID: number; title: string; image: string }) => [
                f.ID,
                { title: f.title, image: f.image },
            ],
        ),
    );

    return {
        favorites,
    };
};
