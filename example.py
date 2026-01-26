import math

def v2_amount_out(dx, rx, ry, fee=0.003):
    if dx <= 0 or rx <= 0 or ry <= 0:
        return 0.0
    dx_eff = dx * (1.0 - fee)
    return (dx_eff * ry) / (rx + dx_eff)

def arb_profit_x_to_x(dx, A_rx, A_ry, B_rx, B_ry, fee=0.003):
    dy = v2_amount_out(dx, A_rx, A_ry, fee)
    x_out = v2_amount_out(dy, B_ry, B_rx, fee)
    return x_out - dx, dy, x_out

def find_best_dx(A_rx, A_ry, B_rx, B_ry, fee=0.003, max_frac=0.3, iters=80):
    lo = 0.0
    hi = max(1e-18, max_frac * A_rx)
    phi = (1 + math.sqrt(5)) / 2
    invphi = 1 / phi
    x1 = hi - (hi - lo) * invphi
    x2 = lo + (hi - lo) * invphi
    f1, _, _ = arb_profit_x_to_x(x1, A_rx, A_ry, B_rx, B_ry, fee)
    f2, _, _ = arb_profit_x_to_x(x2, A_rx, A_ry, B_rx, B_ry, fee)
    for _ in range(iters):
        if f1 < f2:
            lo = x1
            x1, f1 = x2, f2
            x2 = lo + (hi - lo) * invphi
            f2, _, _ = arb_profit_x_to_x(x2, A_rx, A_ry, B_rx, B_ry, fee)
        else:
            hi = x2
            x2, f2 = x1, f1
            x1 = hi - (hi - lo) * invphi
            f1, _, _ = arb_profit_x_to_x(x1, A_rx, A_ry, B_rx, B_ry, fee)
    best_dx = (lo + hi) / 2
    best_profit, best_dy, best_x_out = arb_profit_x_to_x(best_dx, A_rx, A_ry, B_rx, B_ry, fee)
    return best_dx, best_profit, best_dy, best_x_out

def arb_best_two_pools(A_rx, A_ry, B_rx, B_ry, fee=0.003):
    dx1, p1, dy1, xout1 = find_best_dx(A_rx, A_ry, B_rx, B_ry, fee)
    dx2, p2, dy2, xout2 = find_best_dx(B_rx, B_ry, A_rx, A_ry, fee)
    if p1 >= p2:
        return {"route": "A->B", "dx": dx1, "profit": p1, "mid": dy1, "out": xout1}
    return {"route": "B->A", "dx": dx2, "profit": p2, "mid": dy2, "out": xout2}

def main():
    A_rx = 1000.0
    A_ry = 2_000_000.0
    B_rx = 1200.0
    B_ry = 2_000_000.0
    fee = 0.003
    res = arb_best_two_pools(A_rx, A_ry, B_rx, B_ry, fee)
    for k, v in res.items():
        print(f"{k}: {v}")

if __name__ == "__main__":
    main()
