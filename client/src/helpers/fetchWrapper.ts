type WrapperParams = {
  method: "GET" | "POST" | "PUT" | "DELETE";
  path: string;
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
        headers: {
          "Content-Type": "application/json",
        },
        body: params.body,
      },
    );
    const data: Response = await response.json();
    return data;
  } catch (error) {
    console.log(error);
    return {
      status: 500,
      message: "",
      data: [],
      repoInfo: {
        repoName: "",
        ownerName: "",
      },
      reviewed: false,
    };
  }
}
