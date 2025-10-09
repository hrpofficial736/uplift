type WrapperParams = {
  method: "GET" | "POST" | "PUT" | "DELETE";
  path: string;
  body: BodyInit;
};

export default async function callAPI(params: WrapperParams): Promise<object> {
  const response = await fetch(import.meta.env.VITE_SERVER_URL, {
    method: params.method,
    headers: {
      "Content-Type": "application/json",
    },
    body: params.body,
  });
  const data: object = await response.json();
  return data;
}
