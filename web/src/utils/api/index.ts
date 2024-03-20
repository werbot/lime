import { start, done } from "nprogress";

export async function apiGet(url: string, params: object): Promise<any> {
  return handleRequest(url, params, createOptions("GET"));
}

export async function apiPost(
  url: string,
  params: object,
  body: any,
): Promise<any> {
  return handleRequest(url, params, createOptions("POST", body));
}

export async function apiUpdate(
  url: string,
  params: object,
  body: any,
): Promise<any> {
  return handleRequest(url, params, createOptions("PATCH", body));
}

export async function apiDelete(url: string, params: object): Promise<any> {
  return handleRequest(url, params, createOptions("DELETE"));
}

async function handleRequest(
  url: string,
  params: object,
  options: RequestInit,
): Promise<any> {
  start();
  const queryString = new URLSearchParams(
    params as Record<string, string>,
  ).toString();
  const fullUrl = queryString ? `${url}?${queryString}` : url;

  try {
    const response = await fetch(fullUrl, options);
    if (response.status === 401) {
      redirectToSignIn(url);
    }
    return response.json();
//  } catch (error) {
//    console.log(error);
  } finally {
    done();
  }
}

function createOptions(method: string, body?: any): RequestInit {
  const headers: HeadersInit = {};
  let hasContent = false;

  if (body && Object.keys(body).length > 0) {
    headers["Content-Type"] = "application/json";
    body = JSON.stringify(body);
    hasContent = true;
  }

  return {
    credentials: "include",
    method,
    ...(hasContent && { body }),
    ...(Object.keys(headers).length > 0 && { headers }),
  };
}

function redirectToSignIn(url: string): void {
  const signinPath = url.startsWith("/_") ? "/_/signin" : "/signin";
  document.location.href = signinPath;
}
