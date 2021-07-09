import {
  Card,
  CardMedia,
  CardContent,
  CardActions,
  CardHeader,
  Typography,
  Button,
} from "@material-ui/core";
import axios from "axios";
import { GetStaticPaths, GetStaticProps, NextPage } from "next";
import Head from "next/head";
import http from "../../../http";
import { Product } from "../../../model";
import Link from "next/link";

interface ProductDetailPageProps {
  product: Product;
}

const ProductDetailPage: NextPage<ProductDetailPageProps> = ({ product }) => {
  // const router = useRouter();

  // // Enquant a página estiver sendo gerada
  // if (router.isFallback) {
  //   return <div>Carregando...</div>;
  // }

  return (
    <div>
      <Head>
        <title>{product.name} - Detalhes do produto</title>
      </Head>
      {/* Para trabalhar com títulos e textos */}
      <Typography component="h1" variant="h3" color="textPrimary" gutterBottom>
        Produtos
      </Typography>
      <Card>
        <CardHeader
          title={product.name.toUpperCase()}
          subtitle={`R$ ${product.price}`}
        />
        <CardActions>
          <Link
            href="/products/[slug]/order"
            as={`/products/${product.slug}/order`}
          >
            <Button size="small" color="primary" component="a">
              Comprar
            </Button>
          </Link>
        </CardActions>
        <CardMedia image={product.image_url} style={{ paddingTop: "56%" }} />
        <CardContent>
          <Typography variant="body2" color="textSecondary" component="p">
            {product.description}
          </Typography>
        </CardContent>
      </Card>
    </div>
  );
};

export default ProductDetailPage;

export const getStaticProps: GetStaticProps<
  ProductDetailPageProps,
  { slug: string }
> = async (context) => {
  const { slug } = context.params!;

  try {
    const { data: product } = await http.get(`products/${slug}`);

    return {
      props: {
        product,
      },
      revalidate: 1 * 60 * 2, // 2 em 2 minutos
    };
  } catch (e) {
    if (axios.isAxiosError(e) && e.response?.status === 404) {
      return { notFound: true };
    }

    throw e;
  }
};

// Como estamos utilizando o esquema de arquivo com nome de parâmetro,
// e também estamos renderizando a página dinamicamente, já que as informações
// dos produtos vão variar, o next fica em dúvida se é para gerar todas as páginas
// em HTML, ou não. Nesse caso, precisamos informar por meio desta função
export const getStaticPaths: GetStaticPaths = async (context) => {
  const { data: products } = await http.get("products");

  const paths = products.map((product: Product) => ({
    params: { slug: product.slug },
  }));

  // a flag fallback indica como o next irá lidar em casos como:
  // temos 50 produtos no momento do build, ele irá gerar as páginas estáticas para os 50,
  // mas e se adicionarmos um novo produto? Ao invés de precisarmos buildar novamente,
  // podemos solucionar o caso indicando o valor correto para lidar com isso.
  // fallback pode ter 3 valores, sendo true, false ou blocking
  // false => não terá novas informações, e caso tenha, eu mesmo irei rodar o build
  // true => podem ter novas informaçoes, e quando houver, gere um novo arquivo estático,
  // em tempo de execução, caso ele não exista.
  // 'blocking' => a página só será carrega depois que o arquivo tiver sido gerado
  return { paths, fallback: "blocking" };
};
