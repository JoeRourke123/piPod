export const mod = function (m: number, n: number) {
    return ((m % n) + n) % n;
};