class Solution {
        private static int H = 0;
    public int[] solution(int n) {
        int max = n * (n + 1) / 2;
        int[] answer = new int[max];
        int i = 0;
        int num = 1;
        for (int l = 0; l < n; l++) {
            for (int m = l; m < n; m++) {
                i = next(i, l);
                answer[i] = num++;
            }
        }
        answer[i] = max;
        return answer;
    }

    private int next(int i, int l) {
        if (H == 0) return H++;
        return switch (l % 3) {
            case 0 -> i + H++;
            case 1 -> i + 1;
            default -> i - H--;
        };
    }
}