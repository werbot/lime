export function getCookie(cname: string): string {
  const name = cname + "=";
  const decodedCookie = decodeURIComponent(document.cookie);
  const ca = decodedCookie.split(";");

  for (const cookie of ca) {
    let c = cookie.trim();
    if (c.startsWith(name)) {
      return c.substring(name.length);
    }
  }
  return "";
}

export function setCookie(cname: string, cvalue: string, days: number): void {
  const date = new Date();
  date.setDate(date.getDate() + days);
  const value = `${cvalue}${days === null ? "" : "; expires=" + date.toUTCString()}`;
  document.cookie = `${cname}=${value}`;
}

export function getUser(): any {
  const obj = JSON.parse(getCookie("loginUser") || "");
  return obj;
}

export function setUser(user: string): void {
  if (!user) {
    return;
  }

  const value = JSON.stringify(user);
  setCookie("loginUser", value, 1);
}

export function delUser(): void {
  setCookie("loginUser", "", -1);
}
