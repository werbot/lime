import { start, done } from "nprogress";

export async function apiGet(url: string): Promise<any> {
  return handleRequest(url, {
    credentials: "include",
    method: "GET",
  });
}

export async function apiPost(url: string, body: any): Promise<any> {
  const options = createOptions("POST", body);
  return handleRequest(url, options);
}

export async function apiUpdate(url: string, body: any): Promise<any> {
  const options = createOptions("PATCH", body);
  return handleRequest(url, options);
}

export async function apiDelete(url: string): Promise<any> {
  return handleRequest(url, {
    credentials: "include",
    method: "DELETE",
  });
}

async function handleRequest(url: string, options: object): Promise<any> {
  try {
    start();
    const response = await fetch(url, options);
    return response.json();
  } catch (error) {
    console.error(error);
  } finally {
    done();
  }
}

function createOptions(method: string, body: any) {
  const options: {
    credentials: string;
    method: string;
    body?: any;
    headers?: { "Content-Type": string };
  } = {
    credentials: "include",
    method,
  };

  if (body) {
    if (Object.keys(body).length > 0) {
      options.body = JSON.stringify(body);
      options.headers = {
        "Content-Type": "application/json",
      };
    } else {
      options.body = body;
    }
  }
  return options;
}
