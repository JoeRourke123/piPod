import {Params} from "react-router-dom";

export const areMatching = (x: any, y: any): boolean => {
    const ok = Object.keys, tx = typeof x, ty = typeof y;
    return x && y && tx === 'object' && tx === ty ? (
        ok(x).length === ok(y).length &&
        ok(x).every(key => areMatching(x[key], y[key]))
    ) : (x === y);
}

export const getApiUrl = (currentPath: string, params: Readonly<Params> = {}): string => {
    const query = Object.keys(params).map(key => key + '=' + params[key]).join('&');
    return `http://localhost:9091${currentPath}`;
}

export const stringHash = (str: string, seed = 0) => {
    let h1 = 0xdeadbeef ^ seed, h2 = 0x41c6ce57 ^ seed;
    for(let i = 0, ch; i < str.length; i++) {
        ch = str.charCodeAt(i);
        h1 = Math.imul(h1 ^ ch, 2654435761);
        h2 = Math.imul(h2 ^ ch, 1597334677);
    }
    h1  = Math.imul(h1 ^ (h1 >>> 16), 2246822507);
    h1 ^= Math.imul(h2 ^ (h2 >>> 13), 3266489909);
    h2  = Math.imul(h2 ^ (h2 >>> 16), 2246822507);
    h2 ^= Math.imul(h1 ^ (h1 >>> 13), 3266489909);
    // For a single 53-bit numeric return value we could return
    // 4294967296 * (2097151 & h2) + (h1 >>> 0);
    // but we instead return the full 64-bit value:
     const [h3, h4] = [h2>>>0, h1>>>0];
     return h4.toString(36).padStart(7, '0') + h3.toString(36).padStart(7, '0');
};

export const convertMsToTime = (ms: number): string => {
    ms = Math.max(0, ms);
    let minutes = Math.floor(ms / 60000);
    let seconds = ((ms % 60000) / 1000);
    return minutes + ":" + seconds.toFixed(0).padStart(2, '0');
}