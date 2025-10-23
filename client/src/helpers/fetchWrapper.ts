type WrapperParams = {
  method: "GET" | "POST" | "PUT" | "DELETE";
  path: string;
  headers: HeadersInit;
  body: BodyInit;
};

type Response = {
  status: number;
  message: string;
  data: {
    data: {
      text: string;
      toolCalls: null;
    };
    agent: string;
  }[];
  repoInfo: {
    ownerName: string;
    repoName: string;
  };
  reviewed: boolean;
};

export default async function callAPI(
  params: WrapperParams,
): Promise<Response> {
  try {
    const response = await fetch(
      import.meta.env.VITE_SERVER_URL + params.path,
      {
        method: params.method,
        headers: params.headers,
        body: params.body,
      },
    );
    const data: Response = await response.json();
    return data;
  } catch (error) {
    return {
      status: 500,

      message: `${error}`,
      data: [],
      repoInfo: {
        repoName: "",
        ownerName: "",
      },
      reviewed: false,
    };
  }
}
