// Code generated by fastssz. DO NOT EDIT.
// Hash: 2c62e4f3a52cd6b3a37d9953ebbc566ecb0f1f591f68f1f60aa6f5ca0f2bb4a6
// Version: 0.1.3
package testcases

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the Case6 object
func (c *Case6) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(c)
}

// MarshalSSZTo ssz marshals the Case6 object to a target array
func (c *Case6) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'A'
	dst = append(dst, c.A[:]...)

	return
}

// UnmarshalSSZ ssz unmarshals the Case6 object
func (c *Case6) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 32 {
		return ssz.ErrSize
	}

	// Field (0) 'A'
	copy(c.A[:], buf[0:32])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Case6 object
func (c *Case6) SizeSSZ() (size int) {
	size = 32
	return
}

// HashTreeRoot ssz hashes the Case6 object
func (c *Case6) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(c)
}

// HashTreeRootWith ssz hashes the Case6 object with a hasher
func (c *Case6) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'A'
	hh.PutBytes(c.A[:])

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Case6 object
func (c *Case6) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(c)
}