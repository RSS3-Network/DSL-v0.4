#!/usr/bin/python3
# coding=UTF-8

import datetime

import sqlalchemy
from sqlalchemy.dialects.postgresql import BIGINT, TIMESTAMP, TEXT, BYTEA, BOOLEAN, ENUM
from sqlalchemy import Column, Index, create_engine, insert
from sqlalchemy.orm import Session, declarative_base
from enum import Enum
from sqlalchemy.sql import text
import json
import redis
import logging

# dev redis
cache = redis.Redis(
    host='',
    port=6379,
    password='',
    db=1,
)

cache.ping()

# dev
engine = create_engine("xx://xx:xx@xx/xx",
                       echo=True)


Base = declarative_base()


class Network(Enum):
    NetworkNull = ''
    NetworkEthereum = 'ethereum'
    NetworkEthereumClassic = 'ethereum_classic'
    NetworkBinanceSmartChain = 'binance_smart_chain'
    NetworkPolygon = 'polygon'
    NetworkZkSync = 'zksync'
    NetworkXDAI = 'xdai'
    NetworkArweave = 'arweave'
    NetworkArbitrum = 'arbitrum'
    NetworkOptimism = 'optimism'
    NetworkFantom = 'fantom'
    NetworkCelo = 'celo'
    NetworkAvalanche = 'avalanche'
    NetworkCrossbell = 'crossbell'
    NetworkEIP1577 = 'EIP-1577'
    NetworkFarcaster = 'farcaster'
    NetworkAptos = 'aptos'


class Platform(Enum):
    PlatformEmpty = ''
    PlatformMirror = 'Mirror'
    PlatformCrossbell = 'Crossbell'
    PlatformCrossbellXLog = 'xLog'
    PlatformCrossbellXCast = 'xCast'
    PlatformCrossbellXSync = 'xSync'
    PlatformCrossSync = 'CrossSync'
    # PlatformCrossbellio = 'crossbell.io'
    # Platformcrossbell = 'crossbell'
    PlatformFarcaster = 'Farcaster'
    PlatformIQWiki = 'IQ.Wiki'
    PlatformLens = 'Lens'
    PlatformLensLenster = 'Lenster'
    PlatformLensLenstube = 'Lenstube'
    PlatformLensLenstubeBytes = 'lenstube-bytes'
    PlatformLensLensterCrowdfund = 'Lenster Crowdfund'
    PlatformMatters = 'Matters'
    PlatformPOAP = 'POAP'
    PlatformGalaxy = 'Galaxy'
    PlatformEns = 'ENS Registrar'
    PlatformSpaceID = 'Space ID'
    PlatformUnstoppableDomain = 'Unstoppable'
    PlatformAvvy = 'Avvy'
    PlatformSound = 'Sound'
    PlatformGitcoin = 'Gitcoin'
    PlatformSnapshot = 'Snapshot'
    PlatformUniswap = 'Uniswap'
    PlatformSushiswap = 'SushiSwap'
    PlatformPancakeswap = 'PancakeSwap'
    Platform1inch = '1inch'
    PlatformMetamask = 'MetaMask'
    Platform0x = '0x'
    PlatformAAVE = 'AAVE'
    PlatformCurve = 'Curve'
    PlatformLido = 'Lido'
    PlatformPolygonStaking = 'Polygon Staking'
    PlatformTraderJoe = 'TraderJoe'
    PlatformRainbow = 'Rainbow'
    PlatformQuickSwap = 'QuickSwap'
    PlatformKyberSwap = 'KyberSwap'
    PlatformSpookySwap = 'SpookySwap'
    PlatformParaswap = 'Paraswap'
    PlatformDODO = 'DODO'
    PlatformBalancer = 'Balancer'
    PlatformVelodrome = 'Velodrome'
    PlatformOpenSea = 'OpenSea'
    PlatformGem = 'Gem'
    PlatformQuix = 'Quix'
    PlatformLooksRare = 'LooksRare'
    PlatformTofuNFT = 'tofuNFT'
    PlatformBlur = 'Blur'
    PlatformElement = 'Element'
    PlatformHop = 'Hop'
    PlatformCow = 'Cow'
    PlatformMaskNetwork = 'MaskNetwork'
    PlatformMars4 = 'Mars4'
    PlatformAavegotchi = 'Aavegotchi'
    PlatformCarv = 'Carv'
    PlatformPlanetIX = 'PlanetIX'
    PlatformPlanet = 'Planet'
    PlatformCoinbase = 'Coinbase'
    PlatformPolygonBridge = 'Polygon Bridge'
    PlatformBinance = 'Binance'
    PlatformMEXC = 'MEXC: Mexc.com'
    PlatformKuCoin = 'KuCoin'
    PlatformOptimismBridge = 'Optimism Bridge'
    PlatformGate = 'Gate.io'
    PlatformdYdX = 'dYdX'
    PlatformFTXExchange = 'FTX Exchange'
    PlatformOKEx = 'OKEx'
    PlatformHuobi = 'Huobi'
    PlatformCrypto = 'Crypto.com'
    PlatformKucoin = 'Kucoin'
    PlatformKraken = 'Kraken'
    PlatformHotbit = 'Hotbit'
    PlatformBitstamp = 'Bitstamp'
    PlatformPoloniex = 'Poloniex'
    PlatformAscendEX = 'AscendEX'
    PlatformMexc = 'mexc.com'
    PlatformBitMEX = 'BitMEX'
    PlatformBitfinex = 'Bitfinex'
    PlatformLyopay = 'Lyopay'
    PlatformGemini = 'Gemini'
    PlatformBittrex = 'Bittrex'
    PlatformLiqui = 'Liqui.io'
    PlatformBitmart = 'Bitmart'
    PlatformABCC = 'ABCC'
    PlatformAlcumexExchange = 'Alcumex Exchange'
    PlatformAPROBIT = 'APROBIT'
    PlatformArtis = 'Artis'
    PlatformArzPaya = 'ArzPaya.com'
    PlatformATAIX = 'ATAIX'
    PlatformAzbit = 'Azbit'


class Source(Enum):
    SourceNull = ''
    SourceEthereum = 'ethereum'
    SourceOrigin = 'origin'
    SourcePregodETL = 'pregod_etl'
    SourceKurora = 'kurora'
    SourceAlchemy = 'alchemy'
    SourceBlockscout = 'blockscout'
    SourceEIP1577 = 'EIP-1577'
    SourceMoralis = 'moralis'
    SourceZksync = 'zksync'
    SourceAptos = 'aptos'
    SourceSnapshot = 'Snapshot'


class Tag(Enum):
    TagTransaction = 'transaction'
    TagExchange = 'exchange'
    TagCollectible = 'collectible'
    TagSocial = 'social'
    TagDonation = 'donation'
    TagGovernance = 'governance'
    TagMetaverse = 'metaverse'


class Type(Enum):
    Transfer = 'transfer'
    Mint = 'mint'
    Burn = 'burn'
    Approval = 'approval'
    Withdraw = 'withdraw'
    Deposit = 'deposit'
    Trade = 'trade'
    TransactionBridge = 'bridge'
    ExchangeSwap = 'swap'
    ExchangeLiquidity = 'liquidity'
    ExchangeLiquidityAdd = 'add'
    ExchangeLiquidityRemove = 'remove'
    ExchangeLiquidityCollect = 'collect'
    ExchangeLiquiditySupply = 'supply'
    ExchangeLiquidityBorrow = 'borrow'
    ExchangeLiquidityRepay = 'repay'
    CollectiblePoap = 'poap'
    CollectibleMusic = 'music'
    CollectibleMusicBuyEdition = 'buy'
    CollectibleMusicRelease = 'release'
    SocialPost = 'post'
    SocialRevise = 'revise'
    SocialComment = 'comment'
    SocialShare = 'share'
    SocialProfile = 'profile'
    SocialFollow = 'follow'
    SocialUnfollow = 'unfollow'
    SocialLike = 'like'
    SocialWiki = 'wiki'
    SocialReward = 'reward'
    SocialProxy = 'proxy'
    GovernancePropose = 'propose'
    GovernanceVote = 'vote'
    DonationLaunch = 'launch'
    DonationDonate = 'donate'
    MetaverseGift = 'gift'
    MetaverseList = 'list'
    MetaverseUnlist = 'unlist'
    MetaverseClaim = 'claim'


class TransactionsV1(Base):
    __tablename__ = 'transactions_v1'

    block_number = Column("block_number", BIGINT)
    timestamp = Column("timestamp", TIMESTAMP(timezone=True))
    hash = Column("hash", TEXT, primary_key=True, nullable=False)  # primary_key
    index = Column("index", BIGINT, default=0)
    owner = Column("owner", BYTEA, primary_key=True, index=True, nullable=False)  # primary_key, index
    address_from = Column("address_from", BYTEA, index=True, nullable=False)  # index
    address_to = Column("address_to", BYTEA, index=True)  # index
    network = Column("network", ENUM(Network, name="network"), primary_key=True, index=True,
                     nullable=False)  # primary_key, index
    platform = Column("platform", ENUM(Platform, name="platform"), index=True)  # index
    source = Column("source", ENUM(Source, name="source"))
    tag = Column("tag", ENUM(Tag, name="tag"), index=True)  # index
    type = Column("type", ENUM(Type, name="type"), index=True)  # index
    success = Column("success", BOOLEAN)
    fee = Column("fee", TEXT)
    transfers = Column("transfers", BYTEA)
    created_at = Column("created_at", TIMESTAMP(timezone=True), default=datetime.datetime.utcnow,
                        nullable=False)  # index
    updated_at = Column("updated_at", TIMESTAMP(timezone=True), index=True, default=datetime.datetime.utcnow,
                        nullable=False)  # index

    timestamp_index = Index('transactions_v1_timestamp', timestamp.desc())  # index
    index_index = Index('transactions_v1_index', index.desc())  # index
    created_at_index = Index('transactions_v1_created_at', index.desc())  # index


def load_transactions_and_transfers(count=0):
    print("=====start=====")

    # redis cache created_at
    created_at = '2025-11-11 11:11:11'
    if cache.get('transactions_v1') is not None:
        created_at = cache.get('transactions_v1')

    hash_list = ['']
    while len(hash_list) > 0:
        hash_list = []
        hash_map = {}
        transactions_v1 = []

        logging.info("created_at", created_at)

        # read transactions
        with engine.connect() as con:
            query = con.execute(
                text("SELECT * From transactions where created_at <= :v order by created_at desc limit 1000"""),
                {'v': created_at})

        for row in query:
            dict = row._mapping
            platform = dict['platform']
            if dict['platform'] == 'poap':
                platform = 'POAP'
            if dict['platform'] == 'crossbell' or dict['platform'] == 'crossbell.io':
                platform = 'Crossbell'
            if dict['tag'] == '' or dict['type'] == '':
                continue

            hash_list.append(dict['hash'])
            transactions_v1.append(TransactionsV1(
                block_number=dict['block_number'],
                timestamp=dict['timestamp'],
                hash=dict['hash'],
                index=dict['index'],
                owner=bytes(dict['owner'], encoding='utf-8'),
                address_from=bytes(dict['address_from'], encoding='utf-8'),
                address_to=bytes(dict['address_to'], encoding='utf-8'),
                network=dict['network'],
                platform=platform,
                source=dict['source'],
                tag=dict['tag'],
                type=dict['type'],
                success=dict['success'],
                fee=dict['fee'],
                created_at=dict['created_at'],
                updated_at=dict['updated_at'],
            ))

        # read transfers
        query = engine.connect().execute(
            text("""SELECT transaction_hash, tag, type, index, address_from, address_to, metadata, platform, related_urls
            From transfers WHERE transaction_hash in :v"""), {'v': tuple(hash_list)})

        for row in query:
            dict = row._mapping

            transfer = {
                'tag': dict['tag'],
                'type': dict['type'],
                'index': dict['index'],
                'address_from': dict['address_from'],
                'address_to': dict['address_to'],
                'metadata': dict['metadata'],
                'platform': dict['platform'],
                'related_urls': dict['related_urls'],
            }

            if hash_map.get(dict['transaction_hash']) is None:
                hash_map[dict['transaction_hash']] = []

            hash_map[dict['transaction_hash']].append(transfer)

        # deduplicate
        query = engine.connect().execute(
            text("""SELECT hash From transactions_v1 WHERE hash in :v"""),
            {'v': tuple(hash_list), 't': created_at})

        for row in query:
            dict = row._mapping
            hash_map.pop(dict['hash'])

        created_at = transactions_v1[-1].created_at.strftime('%Y-%m-%d %H:%M:%S')
        with Session(engine) as session:
            session.begin()
            platform = {}
            for tx in transactions_v1:
                if hash_map.get(tx.hash) is None:
                    continue

                tx.transfers = bytes(json.dumps(hash_map[tx.hash]), encoding='utf-8')
                platform[tx.platform] = tx.platform
                session.add(tx)
            try:
                session.commit()
                session.close()
            except:
                for k in platform:
                    try:
                        engine.connect().execute(text("""ALTER TYPE platform ADD VALUE :v"""), {'v': k})
                    except sqlalchemy.exc.ProgrammingError:
                        pass
                session.close()
                continue

        cache.set("transactions_v1", created_at)


# create table
Base.metadata.create_all(engine)

load_transactions_and_transfers()
