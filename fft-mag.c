#include <math.h>
#include <stdlib.h>

#include <libavcodec/avfft.h>

#define M_PI 3.14159265358979323846264338327

int main() {
    int N = 1000000; // 1M
    int nbits = 11;
    int input_size = 1 << nbits;
    int output_size = (1 << (nbits - 1)) + 1;

    float *input = malloc(input_size * sizeof(float));
    float *output = malloc(output_size * sizeof(float));

    struct RDFTContext *cx = av_rdft_init(nbits, DFT_R2C);

    float f = M_PI;
    for (int i = 0; i < input_size; ++i) {
        f = floorf(f * M_PI);
        input[i] = f;
    }

    for (int k = 0; k < N; k++ ) {
        av_rdft_calc(cx, input);

        // Calculate magnitudes.
        int n = input_size;
        float n2 = n * n;
        output[0] = 10.0f * log10f(input[0] * input[0] / n2);
        output[n / 2] = 10.0f * log10f(input[1] * input[1] / n2);
        for (int i = 1; i < n / 2; i++) {
            float re = input[i * 2];
            float im = input[i * 2 + 1];
            output[i] = 10.0f * log10f((re * re + im * im) / n2);
        }
    }

    av_rdft_end(cx);
    return 0;
}
