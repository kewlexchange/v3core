import math

def calculate_optimal_dx(r_in1, r_out1, r_in2, r_out2, fee=0.003):
    diff = 1.0 - fee    
    A = (diff**2) * r_out1 * r_out2
    B = r_in1 * r_in2
    C = (diff * r_in2) + (diff**2 * r_out1)
    if A <= B:
        return 0.0
    optimal_dx = (math.sqrt(A * B) - B) / C
    return optimal_dx

def v2_amount_out(dx, rx, ry, fee=0.003):
    if dx <= 0: return 0
    diff = 1.0 - fee
    return (dx * diff * ry) / (rx + dx * diff)

def main():
    A_rx = 1000.0
    A_ry = 2_000_000.0
    B_rx = 1200.0
    B_ry = 2_000_000.0
    fee = 0.003

    best_dx = calculate_optimal_dx(A_rx, A_ry, B_ry, B_rx, fee)
    
    if best_dx > 0:
        dy = v2_amount_out(best_dx, A_rx, A_ry, fee)
        x_out = v2_amount_out(dy, B_ry, B_rx, fee)
        profit = x_out - best_dx
        
        print(f"Optimal Giriş (dx): {best_dx:.18f}")
        print(f"Ara Token (dy): {dy:.18f}")
        print(f"Final Çıkış (x_out): {x_out:.18f}")
        print(f"Maksimum Kâr: {profit:.18f}")
    else:
        print("Kârlı bir arbitraj fırsatı bulunamadı.")

if __name__ == "__main__":
    main()