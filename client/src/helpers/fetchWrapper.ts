export default async function requestServer(
  url: string,
  method: "GET" | "POST" | "PUT" | "DELETE",
  headers: {},
  body: BodyInit
): Promise<{}> {
    const response = await fetch(url, {
        method,
        headers,
        body
    })

    const data = await response.json();
    return data;
}
