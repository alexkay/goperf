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
    }

    av_rdft_end(cx);
    return 0;
}
