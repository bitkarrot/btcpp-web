package helpers

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestInvoiceFallbacks(t *testing.T) {
	tests := []struct {
		invoice string
		result  []string
	}{{
		invoice: "lnbc20m1pvjluezsp5zyg3zyg3zyg3zyg3zyg3zyg3zyg3zyg3zyg3zyg3zyg3zyg3zygshp58yjmdan79s6qqdhdzgynm4zwqd5d7xmw5fk98klysy043l2ahrqspp5qqqsyqcyq5rqwzqfqqqsyqcyq5rqwzqfqqqsyqcyq5rqwzqfqypqfppqw508d6qejxtdg4y5r3zarvary0c5xw7k9qrsgqt29a0wturnys2hhxpner2e3plp6jyj8qx7548zr2z7ptgjjc7hljm98xhjym0dg52sdrvqamxdezkmqg4gdrvwwnf0kv2jdfnl4xatsqmrnsse",
		result:  []string{"bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t4"},
	}, {
		/* Note: the networks don't match */
		invoice: "lnbcrt1068340n1pjk5fhasp55zt4kynx8484qug05vcxj362fhkllxk4wa35vqrnmphpz5nyhx6spp5x04x95w7hkvlvyge3s3qm7knkuqmq0c4acx5509t82n79cxgfs3sdxcxys8g6trddjhggrxdaezqargv5sxy6t5vdhkjm3t9vsyyat9dehhxgzpd9ex2ueqxgcrydpvypcxz7tdv4h8gueqv4jxjarfdahzq2pyxsc4256y9ys9ken9xdskxvm9xvkk2c35vyknge3cx5knsetr8ykkvefcxyersdnxvsmrvvrude5kvaredejkjsr8d4skjmpwvdhk6lrvda3kzmzaxqzpucqp2fp4p9fap4n9vtkdc5n5wq3cvqyw2wsdg6yjxkgtwzvt0kk769ez9m70q9qx3qysgqktdh7uxdezphwcrdllqqk3jgjc5lv3htfw3332svcuvkqwgp66mxxrahuesvjemr6y5kqlfg2hu69ea82zezwdmhjq6jv206a6nnqcsqtfv5et",
		result:  []string{"bc1p9fap4n9vtkdc5n5wq3cvqyw2wsdg6yjxkgtwzvt0kk769ez9m70q4dlkk7"},
	}, {
		invoice: "lnbc20m1pvjluezsp5zyg3zyg3zyg3zyg3zyg3zyg3zyg3zyg3zyg3zyg3zyg3zyg3zygshp58yjmdan79s6qqdhdzgynm4zwqd5d7xmw5fk98klysy043l2ahrqspp5qqqsyqcyq5rqwzqfqqqsyqcyq5rqwzqfqqqsyqcyq5rqwzqfqypqfp4qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3q9qrsgq9vlvyj8cqvq6ggvpwd53jncp9nwc47xlrsnenq2zp70fq83qlgesn4u3uyf4tesfkkwwfg3qs54qe426hp3tz7z6sweqdjg05axsrjqp9yrrwc",
		result:  []string{"bc1qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3qccfmv3"},
	}, {
		"lnbc25m1pvjluezpp5qqqsyqcyq5rqwzqfqqqsyqcyq5rqwzqfqqqsyqcyq5rqwzqfqypqdq5vdhkven9v5sxyetpdeessp5zyg3zyg3zyg3zyg3zyg3zyg3zyg3zyg3zyg3zyg3zyg3zyg3zygs9q5sqqqqqqqqqqqqqqqqsgq2qrqqqfppnqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqppnqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqpp4qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqhpnqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqhp4qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqspnqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqsp4qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqnp5qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqnpkqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqz599y53s3ujmcfjp5xrdap68qxymkqphwsexhmhr8wdz5usdzkzrse33chw6dlp3jhuhge9ley7j2ayx36kawe7kmgg8sv5ugdyusdcqzn8z9x",
		[]string{},
	}}
	for i, input := range tests {
		addrs, err := GetFallbackAddrs(input.invoice)
		if err != nil {
			t.Fatalf("index: %d, %s", i, err)
		}

		if !cmp.Equal(addrs, input.result) {
			t.Fatalf("index: %d, exp: %v, got: %v", i, input.result, addrs)
		}
	}
}
