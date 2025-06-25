import java.util.stream.LongStream;

public class Main {
    public static void main(String[] args) {
        long start = System.nanoTime();
        double result = LongStream.range(1, 100_000_001)
                .parallel()
                .mapToDouble(Math::sqrt)
                .sum();
        long end = System.nanoTime();
        System.out.printf("Result: %.2f, Time: %.2f seconds%n", result, (end - start) / 1e9);
    }
}
