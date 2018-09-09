using System;
using System.IO;
using System.IO.Compression;

namespace myApp
{
    class Program
    {
        static void Main(string[] args)
        {
            FileInfo fileToDecompress = new FileInfo("newhello.bin");
            GZipDecompress(fileToDecompress);
            // string inputfilename = "hello.bin";
            // byte[] fileBytes = File.ReadAllBytes(inputfilename);

            // //string inputStr = "";
            // //byte[] inputBytes = Convert.FromBase64String(inputStr);
            
            // using (var inputStream = new MemoryStream(fileBytes))
            // using (var gZipStream = new GZipStream(inputStream, CompressionMode.Decompress))
            // using (var streamReader = new StreamReader(gZipStream))
            // {
            //     var decompressed = streamReader.ReadToEnd();
            
            //     Console.WriteLine(decompressed);
            //     Console.ReadLine();
            // }


            Console.WriteLine("Hello World! Again!");
        }

        public static void Decompress(FileInfo fileToDecompress)
        {
            using (FileStream originalFileStream = fileToDecompress.OpenRead())
            {
                string currentFileName = fileToDecompress.FullName;
                string newFileName = currentFileName.Remove(currentFileName.Length - fileToDecompress.Extension.Length);
                using (FileStream decompressedFileStream = File.Create(newFileName))
                {
                    using (DeflateStream decompressionStream = new DeflateStream(originalFileStream, CompressionMode.Decompress))
                    {
                        decompressionStream.CopyTo(decompressedFileStream);
                        Console.WriteLine("Decompressed: {0}", fileToDecompress.Name);
                    }
                }
            }
        }
        
        public static void GZipDecompress(FileInfo fileToDecompress)
        {
            using (FileStream originalFileStream = fileToDecompress.OpenRead())
            {
                string currentFileName = fileToDecompress.FullName;
                string newFileName = currentFileName.Remove(currentFileName.Length - fileToDecompress.Extension.Length);
                using (FileStream decompressedFileStream = File.Create(newFileName))
                {
                    using (GZipStream decompressionStream = new GZipStream(originalFileStream, CompressionMode.Decompress))
                    {
                        decompressionStream.CopyTo(decompressedFileStream);
                        Console.WriteLine("Decompressed: {0}", fileToDecompress.Name);

                    }
                }
            }
        }

    }
}
