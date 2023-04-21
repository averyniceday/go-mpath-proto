/*
 * Genome Nexus API
 *
 * This page shows how to use HTTP requests to access the Genome Nexus API. There are more high level clients available in Python, R, JavaScript, TypeScript and various other languages as well as a command line client to annotate MAF and VCF. See https://docs.genomenexus.org/api.  Aside from programmatic clients there are web based tools to annotate variants, see https://docs.genomenexus.org/tools.   We currently only provide long-term support for the '/annotation' endpoint. The other endpoints might change.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ClinVar struct {

	// allele_id
	AlleleId int32 `json:"alleleId,omitempty"`

	// alt
	Alt string `json:"alt,omitempty"`

	// chrom
	Chrom string `json:"chrom,omitempty"`

	// cytogenic
	Cytogenic string `json:"cytogenic,omitempty"`

	// gene
	Gene *Gene `json:"gene,omitempty"`

	// hg19
	Hg19 *Hg19 `json:"hg19,omitempty"`

	// hg38
	Hg38 *Hg38 `json:"hg38,omitempty"`

	// hgvs
	Hgvs *Hgvs `json:"hgvs,omitempty"`

	// license
	License string `json:"license,omitempty"`

	Rcv []Rcv `json:"rcv,omitempty"`

	// variant_id
	VariantId int32 `json:"variantId,omitempty"`
}